package event

import (
	
	"net/http"

	getEvent "github.com/sea350/ustart_go/get/event"
	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//ManageEvents ...
func ManageEvents(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	var heads []types.FloatingHead

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	for _, eventInfo := range userstruct.Events {
		var isAdmin = false
		event, err := getEvent.EventByID(client.Eclient, eventInfo.EventID)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}

		for _, memberInfo := range event.Members {
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
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}
		heads = append(heads, head)
	}

	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "ManageEventMembers", cs)

}
