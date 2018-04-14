package project

import (
	"fmt"
	"net/http"
	"time"

	client "github.com/sea350/ustart_go/middleware/client"
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
		fmt.Println("err middleware/project/acceptjoinrequest line 27")
		fmt.Println(err)
	}

	err = userPost.AppendProject(client.Eclient, newMemberID, types.ProjectInfo{ProjectID: projID, Visible: true})
	if err != nil {
		fmt.Println("err middleware/project/acceptjoinrequest line 34")
		fmt.Println(err)
	}

	var newMember types.Member
	newMember.MemberID = newMemberID
	newMember.Role = 2
	newMember.Title = "Member"
	newMember.Visible = true
	newMember.JoinDate = time.Now()

	err = projPost.AppendMember(client.Eclient, projID, newMember)
	if err != nil {
		fmt.Println("err middleware/project/acceptjoinrequest line 48")
		fmt.Println(err)
	}

	fmt.Fprintln(w, newNumRequests)
}
