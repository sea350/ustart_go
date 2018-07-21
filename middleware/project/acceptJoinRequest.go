package project

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	getChat "github.com/sea350/ustart_go/get/chat"
	getProj "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//AcceptJoinRequest ...
func AcceptJoinRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projID := r.FormValue("projectID")
	newMemberID := r.FormValue("userID")

	newNumRequests, err := uses.RemoveRequest(client.Eclient, projID, newMemberID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	err = userPost.AppendProject(client.Eclient, newMemberID, types.ProjectInfo{ProjectID: projID, Visible: true})
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	theRole := r.FormValue("role")

	var newMember types.Member
	switch theRole {
	case "Moderator":
		newMember.Title = "Admin"
		newMember.Role = 1
	default:
		newMember.Title = "Member"
		newMember.Role = 2
	}

	newMember.MemberID = newMemberID
	newMember.Visible = true
	newMember.JoinDate = time.Now()

	err = projPost.AppendMember(client.Eclient, projID, newMember)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return
	}

	proj, err := getProj.ProjectByID(client.Eclient, projID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return
	}

	convo, err := getChat.ConvoByID(client.Eclient, proj.Subchats[0].ConversationID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return
	}

	convo.Eavesdroppers[newMemberID] = types.Eavesdropper{Class: 1}

	err = postChat.UpdateConvo(client.Eclient, proj.Subchats[0].ConversationID, "Eavesdroppers", convo.Eavesdroppers)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return
	}

	proxyID, err := getChat.ProxyIDByUserID(client.Eclient, newMemberID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return
	}

	err = postChat.AppendToProxy(client.Eclient, proxyID, proj.Subchats[0].ConversationID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return
	}

	fmt.Fprintln(w, newNumRequests)
}
