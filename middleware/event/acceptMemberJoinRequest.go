package event

import (
	"fmt"
	"net/http"
	"time"

	client "github.com/sea350/ustart_go/middleware/client"
	evntPost "github.com/sea350/ustart_go/post/event"
	userPost "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//AcceptMemberJoinRequest
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
		fmt.Println("err middleware/event/acceptmemberjoinrequest line 27")
		fmt.Println(err)
	}

	err = userPost.AppendEvent(client.Eclient, newMemberID, types.EventInfo{EventID: evntID, Visible: true})
	if err != nil {
		fmt.Println("err middleware/event/acceptmemberjoinrequest line 33")
		fmt.Println(err)
	}

	var newMember types.EventMembers
	newMember.MemberID = newMemberID
	newMember.Role = 2
	newMember.Title = "Member"
	newMember.Visible = true
	newMember.JoinDate = time.Now()

	err = evntPost.AppendMember(client.Eclient, evntID, newMember)
	if err != nil {
		fmt.Println("err middleware/event/acceptmemberjoinrequest line 46")
		fmt.Println(err)
	}

	fmt.Fprintln(w, newNumRequests)
}
