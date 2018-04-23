package project

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	get "github.com/sea350/ustart_go/get/project"
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A PROFILE VIEW REQUEST")
	session, _ := store.Get(r, "session_please")
	test1, boo := session.Values["DocID"]

	fmt.Println("LINE 26 bool", boo)

	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

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

	Proj, err = get.ProjectByID(eclient, data.ProjectID)
	if err != nil {
		fmt.Println("error line 75 profile/handler.go")
	}

	resp.update(false, errors.New(" This is an unknown error"), Proj)
	//	marshalledData, err := json.Marshal(data)

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	if test1 == test1 {
		isLeader, index := uses.IsLeader(eclient, data.ProjectID, data.SessUser.DocID)
		switch data.Intent {
		case "join":
			if data.Username == data.Username {
				//isMember := get.IsMember(eclient, data.SessUser.DocID, Proj) //session.Values["DocID"].(string))

				if index == -1 {
					fmt.Println("INTENT TO JOIN")
					err1 := userPost.AppendSentProjReq(eclient, data.SessUser.DocID, data.ProjectID)

					err2 := projPost.AppendMemberReqReceived(eclient, data.ProjectID, data.SessUser.DocID)

					if err1 != nil {
						resp.update(err1 == nil, err1, Proj)
					} else {
						resp.update(err2 == nil, err2, Proj)
					}

				} else if index != -1 && !isLeader {
					fmt.Println("INTENT TO LEAVE")
					err = projPost.DeleteMember(eclient, data.ProjectID, data.SessUser.DocID)

					resp.update(err == nil, err, Proj)
				}
			}
		case "accept":
			if isLeader {
				err1 := userPost.AppendProject(eclient, data.SessUser.DocID, types.ProjectInfo{ProjectID: data.ProjectID, Visible: true})
				_, err2 := uses.RemoveRequest(eclient, data.ProjectID, data.JoinerID)

				var newMember types.Member
				newMember.JoinDate = time.Now()
				newMember.MemberID = data.JoinerID
				newMember.Role = 1
				newMember.Title = "NewMem"
				newMember.Visible = true

				err3 := projPost.AppendMember(eclient, data.ProjectID, newMember)

				if err1 != nil {
					resp.update(err1 == nil, err1, Proj)
				} else if err2 != nil {
					resp.update(err2 == nil, err2, Proj)
				} else {
					resp.update(err3 == nil, err3, Proj)
				}

			}
		case "reject":
			_, err = uses.RemoveRequest(eclient, data.ProjectID, data.JoinerID)
			resp.update(err == nil, err, Proj)

		}
	} else {
		resp.update(false, errors.New("Something went wrong"), Proj)
	}

}
