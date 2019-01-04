package widget

import (
	"net/http"

	getEvnt "github.com/sea350/ustart_go/get/event"

	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/widget"
	"github.com/sea350/ustart_go/uses"
)

//AddEventWidget ... After widget form submission adds a widget to database
func AddEventWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	evnt, member, err := getEvnt.EventAndMember(client.Eclient, r.FormValue("eventWidget"), test1.(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	if uses.HasEventPrivilege("widget", evnt.PrivilegeProfiles, member) {

		/* empty eventWidget */
		if r.FormValue("eventWidget") == `` {

			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Event ID Missing")
			http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
			return
		}

		newWidget, err := ProcessWidgetForm(r)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
		}

		newWidget.UserID = r.FormValue("eventWidget")
		if r.FormValue("editID") == `0` {
			err := uses.AddWidget(client.Eclient, r.FormValue("eventWidget"), newWidget, false, true)
			if err != nil {
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			}
		} else {
			err := post.ReindexWidget(client.Eclient, r.FormValue("editID"), newWidget)
			if err != nil {
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			}
		}
		http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
	} else {

		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "You do not have the privilege to add a widget to this event. Check your privilege. ")
		http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
	}
}
