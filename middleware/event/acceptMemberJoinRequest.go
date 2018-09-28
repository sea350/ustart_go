package event

import (
	"fmt"
	"log"
	"net/http"
	"time"

	client "github.com/sea350/ustart_go/middleware/client"
	evntPost "github.com/sea350/ustart_go/post/event"
	userPost "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//AcceptMemberJoinRequest ...
func AcceptMemberJoinRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	evntID := r.FormValue("eventID")
	newMemberID := r.FormValue("userID")

	newNumRequests, err := uses.RemoveEventRequest(client.Eclient, evntID, newMemberID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	err = userPost.AppendEvent(client.Eclient, newMemberID, types.EventInfo{EventID: evntID, Visible: true})
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	theRole := r.FormValue("role")
	var newMember types.EventMembers
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

	err = evntPost.AppendMember(client.Eclient, evntID, newMember)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, newNumRequests)
}
