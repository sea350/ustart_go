package event

import (
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ViewEvent ... rendering the event
//ProjectsPage
func ViewEvent(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	event, err := uses.AggregateEventData(client.Eclient, r.URL.Path[7:], test1.(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}
	widgets, errs := uses.LoadWidgets(client.Eclient, event.EventData.Widgets)
	if len(errs) > 0 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "there were one or more errors loading widgets")
		for _, eror := range errs {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", eror)
		}
	}

	userstruct, err := get.UserByID(client.Eclient, test1.(string))
	if err != nil {
		panic(err)
	}

	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), Event: event, Widgets: widgets}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "events", cs)

}
