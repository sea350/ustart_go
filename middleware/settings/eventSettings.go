package settings

import (
	
	"net/http"
	

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Event ...
func Event(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	evntURL := r.URL.Path[15:]
	event, err := uses.AggregateEventData(client.Eclient, evntURL, test1.(string))
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	var isAdmin = false
	for _, member := range event.EventData.Members {
		if member.MemberID == test1.(string) && member.Role <= 0 {
			isAdmin = true
			break
		}
	}
	if isAdmin {
		cs := client.ClientSide{Event: event}
		client.RenderSidebar(w, r, "template2-nil")
		client.RenderSidebar(w, r, "leftnav-nil")
		client.RenderTemplate(w, r, "eventSettings", cs)

	} else {
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

}
