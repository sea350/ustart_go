package event

import (
	"fmt"
	"net/http"

	getEvent "github.com/sea350/ustart_go/get/event"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/uses"

	"github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
)

//ManageEvents ...
func ManageEvents(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	var heads []types.FloatingHead

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		fmt.Println(err)
		fmt.Println("err: middleware/event/manageevents Line 26")
	}

	for _, eventInfo := range userstruct.Events {
		if eventInfo.EventID == nil {
			fmt.Println("Missing EventID from %s", userstruct.Username)
			continue
		}
		var isAdmin = false
		evnt, err := getEvent.EventByID(client.Eclient, eventInfo.EventID)
		if err != nil {
			fmt.Println(err)
			fmt.Println("err: middleware/event/manageevent Line 35")
		}
		for _, memberInfo := range evnt.Members {
			if memberInfo.MemberID == test1.(string) && memberInfo.Role <= 1 {
				//finds user in the list of members and also checks if they have creator rank
				isAdmin = true
				//head.Followed in this case expresses whether or not they have edit permissions
			}
		}
		if !isAdmin {
			continue
		}
		head, err := uses.ConvertEventToFloatingHead(client.Eclient, eventInfo.EventID)
		if err != nil {
			fmt.Println(err)
			fmt.Println("err: middleware/event/manageevents Line 51")
		}
		heads = append(heads, head)
	}

	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "eventManager", cs)
}
