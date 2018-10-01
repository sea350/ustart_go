package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	get "github.com/sea350/ustart_go/backend/get/event"
	evntPost "github.com/sea350/ustart_go/backend/post/event"
	userPost "github.com/sea350/ustart_go/backend/post/user"
	"github.com/sea350/ustart_go/backend/types"
	uses "github.com/sea350/ustart_go/backend/uses"

	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

//Handler responds to http requests about content
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING AN EVEN VIEW REQUEST")

	Evnt := types.Events{}
	//Steup the response
	resp := &response{
		Successful: false,
		Event:      Evnt,
		ErrMsg:     errors.New("Unknown failure"),
	}

	if acrh, ok := r.Header["Access-Control-Request-Headers"]; ok {
		w.Header().Set("Access-Control-Allow-Headers", acrh[0])
	}
	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if acao, ok := r.Header["Access-Control-Allow-Origin"]; ok {
		w.Header().Set("Access-Control-Allow-Origin", acao[0])
	} else {
		if _, oko := r.Header["Origin"]; oko {
			w.Header().Set("Access-Control-Allow-Origin", r.Header["Origin"][0])
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Connection", "Close")

	defer json.NewEncoder(w).Encode(resp)

	// Parse the request.
	fmt.Println("Parsing request")

	//var sessUsr types.AppSessionUser
	data := form{}

	err := json.NewDecoder(r.Body).Decode(&data)

	fmt.Println("THIS IS THE DATA:", data)
	if err != nil {
		fmt.Println("error line 70 profile/handler.go")
	}

	evntID, err := get.EventIDByURL(eclient, data.Username)

	fmt.Println("THE EVENT ID SHOULD BE:", evntID)
	if err != nil {
		fmt.Println("error line 79 profile/handler.go")
	}

	Evnt, err = get.EventByID(eclient, evntID)
	if err != nil {
		fmt.Println("error line 83 profile/handler.go")
	}

	//resp.update(false, errors.New(" This is an unknown error"), Proj)
	//	marshalledData, err := json.Marshal(data)

	if 1 == 1 {
		fmt.Println("EVENT INFO", Evnt)
		isLeader, index := uses.IsEventLeader(eclient, evntID, data.SessUser.DocID)
		fmt.Println("THE INDEX IS:", index)
		switch data.Intent {
		case "member join":
			if data.Username == data.Username {
				isMember := get.IsEventMember(eclient, data.SessUser.DocID, Evnt) //session.Values["DocID"].(string))
				if !isMember {
					fmt.Println("INTENT TO JOIN")
					err1 := userPost.AppendSentEventReq(eclient, data.SessUser.DocID, evntID)

					err2 := evntPost.AppendMemberReqReceived(eclient, evntID, data.SessUser.DocID)
					if err1 != nil {
						resp.update(err1 == nil, err1, Evnt)
					} else {
						resp.update(err2 == nil, err2, Evnt)
					}
					resp.update(err2 == nil, err2, Evnt)

				} else if index != -1 && !isLeader {
					fmt.Println("INTENT TO LEAVE")
					err = evntPost.DeleteMember(eclient, evntID, data.SessUser.DocID)
					resp.update(err == nil, err, Evnt)
				}
			}
		case "guest join":
			if data.Username == data.Username {
				isGuest := get.IsEventGuest(eclient, data.SessUser.DocID, Evnt) //session.Values["DocID"].(string))
				if !isGuest {
					fmt.Println("INTENT TO JOIN")
					err1 := userPost.AppendSentEventReq(eclient, data.SessUser.DocID, evntID)

					classification, err := strconv.Atoi(r.FormValue("classification"))
					if err != nil {
						resp.update(err == nil, err, Evnt)
					}
					err2 := evntPost.AppendGuestReqReceived(eclient, evntID, data.SessUser.DocID, classification)
					if err1 != nil {
						resp.update(err1 == nil, err1, Evnt)
					} else {
						resp.update(err2 == nil, err2, Evnt)
					}
					resp.update(err2 == nil, err2, Evnt)

				} else if index != -1 && !isLeader {
					fmt.Println("INTENT TO LEAVE")
					err = evntPost.DeleteGuest(eclient, evntID, data.SessUser.DocID)
					resp.update(err == nil, err, Evnt)
				}
			}
		case "accept member":
			if isLeader {
				err1 := userPost.AppendEvent(eclient, data.MemberJoinerID, types.EventInfo{EventID: evntID, Visible: true})
				_, err2 := uses.RemoveEventRequest(eclient, evntID, data.MemberJoinerID)

				var newMember types.EventMembers
				newMember.JoinDate = time.Now()
				newMember.MemberID = data.MemberJoinerID
				newMember.Role = 1
				newMember.Title = "NewMem"
				newMember.Visible = true

				err3 := evntPost.AppendMember(eclient, evntID, newMember)
				if err1 != nil {
					resp.update(err1 == nil, err1, Evnt)
				} else if err2 != nil {
					resp.update(err2 == nil, err2, Evnt)
				} else {
					resp.update(err3 == nil, err3, Evnt)
				}
			}

		case "accept guest":
			err1 := userPost.AppendEvent(eclient, data.GuestJoinerID, types.EventInfo{EventID: evntID, Visible: true})
			classification, err := strconv.Atoi(r.FormValue("classification"))
			if err != nil {
				resp.update(err == nil, err, Evnt)
			}
			_, err2 := uses.RemoveGuestRequest(eclient, evntID, data.GuestJoinerID, classification)

			var newGuest types.EventGuests
			newGuest.GuestID = data.MemberJoinerID
			newGuest.Status = 0
			// newGuest.Representative = data.Representative
			// newGuest.Visible = true

			err3 := evntPost.AppendGuest(eclient, evntID, newGuest)

			if err1 != nil {
				resp.update(err1 == nil, err1, Evnt)
			} else if err2 != nil {
				resp.update(err2 == nil, err2, Evnt)
			} else {
				resp.update(err3 == nil, err3, Evnt)
			}

		case "reject member":
			_, err = uses.RemoveEventRequest(eclient, evntID, data.MemberJoinerID)
			resp.update(err == nil, err, Evnt)
		case "reject guest":
			classification, err := strconv.Atoi(r.FormValue("classification"))
			if err != nil {
				resp.update(err == nil, err, Evnt)
			}
			_, err = uses.RemoveGuestRequest(eclient, evntID, data.GuestJoinerID, classification)
			resp.update(err == nil, err, Evnt)

		case "member leave":
			if isLeader {
				err1 := uses.NewEventLeader(eclient, evntID, data.SessUser.DocID, data.MemberJoinerID)
				err2 := evntPost.DeleteMember(eclient, evntID, data.SessUser.DocID)
				if err1 != nil {
					resp.update(err2 == nil, err2, Evnt)
				} else {
					resp.update(err2 == nil, err2, Evnt)
				}
			} else {
				err = evntPost.DeleteMember(eclient, evntID, data.SessUser.DocID)
			}

		case "guest leave":
			err = evntPost.DeleteGuest(eclient, evntID, data.SessUser.DocID)

		case "get":
			Evnt, errGet := get.EventByID(eclient, evntID)
			resp.update(errGet == nil, errGet, Evnt)
		}
	} else {
		resp.update(false, errors.New("Something went wrong"), Evnt)
	}
}
