package project

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
)

//LoadJoinRequests ...
func LoadJoinRequests(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	ID := r.FormValue("UNKNOWN") //projectID
	var heads []types.FloatingHead

	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("err: middleware/project/loadjoinrequest Line 26")
	}

	for index, userID := range proj.MemberReqReceived {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, userID)
		if err != nil {
			fmt.Println(err)
			fmt.Println(fmt.Sprintf("err: middleware/project/loadjoinrequest, Line 35, index %d", index))
		}
		heads = append(heads, head)
	}

}