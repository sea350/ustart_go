package profile

import (
	
	"net/http"
	

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//AjaxChangeEventVisibility ... an ajax call that changes whether a project is visible on the user page
func AjaxChangeEventVisibility(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	ID, _ := session.Values["DocID"]
	if ID == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	eventID := r.FormValue("eventID")

	usr, err := get.UserByID(client.Eclient, ID.(string))
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		return
	}

	for i := range usr.Events {
		if usr.Events[i].EventID == eventID {
			if usr.Events[i].Visible {
				usr.Events[i].Visible = false
			} else {
				usr.Events[i].Visible = true
			}
		}
	}

	err = post.UpdateUser(client.Eclient, ID.(string), "Events", usr.Events)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
}
