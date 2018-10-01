package project

import (
	"fmt"
	"log"
	"net/http"
	"time"

	getChat "github.com/sea350/ustart_go/backend/get/chat"
	getProj "github.com/sea350/ustart_go/backend/get/project"
	client "github.com/sea350/ustart_go/backend/middleware/client"
	postChat "github.com/sea350/ustart_go/backend/post/chat"
	post "github.com/sea350/ustart_go/backend/post/notification"
	projPost "github.com/sea350/ustart_go/backend/post/project"
	userPost "github.com/sea350/ustart_go/backend/post/user"
	types "github.com/sea350/ustart_go/backend/types"
	uses "github.com/sea350/ustart_go/backend/uses"
)

//AcceptJoinRequest ... made for ajax
func AcceptJoinRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projID := r.FormValue("projectID")
	if projID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("WARNING: project ID not received")
		return
	}
	newMemberID := r.FormValue("userID")
	if newMemberID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("WARNING: new member ID not received")
		return
	}

	newNumRequests, err := uses.RemoveRequest(client.Eclient, projID, newMemberID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	proj, err := getProj.ProjectByID(client.Eclient, projID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	for i := range proj.Members {
		if proj.Members[i].MemberID == newMemberID {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("USER IS ALREADY A MEMBER")
			return
		}
	}

	err = userPost.AppendProject(client.Eclient, newMemberID, types.ProjectInfo{ProjectID: projID, Visible: true})
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
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
		log.Println(err)
		return
	}

	convo, err := getChat.ConvoByID(client.Eclient, proj.Subchats[0].ConversationID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	convo.Eavesdroppers = append(convo.Eavesdroppers, types.Eavesdropper{Class: 1, DocID: newMemberID})

	err = postChat.UpdateConvo(client.Eclient, proj.Subchats[0].ConversationID, "Eavesdroppers", convo.Eavesdroppers)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	proxyID, err := getChat.ProxyIDByUserID(client.Eclient, newMemberID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	err = postChat.AppendToProxy(client.Eclient, proxyID, proj.Subchats[0].ConversationID, false)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	var newNotif types.Notification
	newNotif.ProjectJoinRequestAccepted(newMemberID, projID)
	_, err = post.IndexNotification(client.Eclient, newNotif)

	fmt.Fprintln(w, newNumRequests)
}
