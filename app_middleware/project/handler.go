package project

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	get "github.com/sea350/ustart_go/get/project"
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

//var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A PROJECT VIEW REQUEST")

	Proj := types.Project{}
	// Setup the response
	resp := &response{
		Successful: false,
		Project:    Proj,

		ErrMsg: errors.New("Unknown failure"),
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

	projID, err := get.ProjectIDByURL(eclient, data.Username)

	fmt.Println("THE PROJECT ID SHOULD BE:", projID)
	if err != nil {
		fmt.Println("error line 79 profile/handler.go")
	}
	Proj, err = get.ProjectByID(eclient, projID)
	if err != nil {
		fmt.Println("error line 83 profile/handler.go")
	}

	//resp.update(false, errors.New(" This is an unknown error"), Proj)
	//	marshalledData, err := json.Marshal(data)

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	if 1 == 1 {
		fmt.Println("PROJECT INFO", Proj)
		isLeader, index := uses.IsLeader(eclient, projID, data.SessUser.DocID)
		fmt.Println("tHE INDEX IS:", index)
		switch data.Intent {
		case "join":
			if data.Username == data.Username {
				isMember := get.IsMember(eclient, data.SessUser.DocID, Proj) //session.Values["DocID"].(string))

				if !isMember {
					fmt.Println("INTENT TO JOIN")
					err1 := userPost.AppendSentProjReq(eclient, data.SessUser.DocID, projID)

					err2 := projPost.AppendMemberReqReceived(eclient, projID, data.SessUser.DocID)

					if err1 != nil {
						resp.update(err1 == nil, err1, Proj)
					} else {
						resp.update(err2 == nil, err2, Proj)
					}

					resp.update(err2 == nil, err2, Proj)
				} else if index != -1 && !isLeader {
					fmt.Println("INTENT TO LEAVE")
					err = projPost.DeleteMember(eclient, projID, data.SessUser.DocID)

					resp.update(err == nil, err, Proj)
				}
			}
		case "accept":
			if isLeader {
				err1 := userPost.AppendProject(eclient, data.JoinerID, types.ProjectInfo{ProjectID: projID, Visible: true})
				_, err2 := uses.RemoveRequest(eclient, projID, data.JoinerID)

				var newMember types.Member
				newMember.JoinDate = time.Now()
				newMember.MemberID = data.JoinerID
				newMember.Role = 1
				newMember.Title = "NewMem"
				newMember.Visible = true

				err3 := projPost.AppendMember(eclient, projID, newMember)

				if err1 != nil {
					resp.update(err1 == nil, err1, Proj)
				} else if err2 != nil {
					resp.update(err2 == nil, err2, Proj)
				} else {
					resp.update(err3 == nil, err3, Proj)
				}

			}
		case "reject":
			_, err = uses.RemoveRequest(eclient, projID, data.JoinerID)
			resp.update(err == nil, err, Proj)

		case "leave":
			if isLeader {
				err1 := uses.NewProjectLeader(eclient, projID, data.SessUser.DocID, data.JoinerID)
				err2 := projPost.DeleteMember(eclient, projID, data.SessUser.DocID)

				if err1 != nil {
					resp.update(err2 == nil, err2, Proj)
				} else {
					resp.update(err2 == nil, err2, Proj)
				}
			} else {
				err = projPost.DeleteMember(eclient, projID, data.SessUser.DocID)

			}
		case "get":
			Proj, errGet := get.ProjectByID(eclient, projID)
			resp.update(errGet == nil, errGet, Proj)

		}
	} else {
		resp.update(false, errors.New("Something went wrong"), Proj)
	}

}
