package settings

import (
	"net/http"

	"github.com/microcosm-cc/bluemonday"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
)

//LeaveEventGuest ... lets a user leave a event
//If Rol
func LeaveEventGuest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()
	leavingUser := p.Sanitize(r.FormValue("leaverID"))
	evntID := p.Sanitize(r.FormValue("eventID"))

	evnt, err := get.EventByID(client.Eclient, evntID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	var canLeave = false
	if leavingUser == test1.(string) {
		//if the current active user wants to leave, they can
		canLeave = true
	} else {
		for _, mem := range evnt.Guests {
			if mem.GuestID == test1.(string) {
				//if the current acessing user is creator, they can do whatever they want
				canLeave = true
				break
			}
		}
	}
	if !canLeave {
		http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
		return
	}

	err = post.DeleteGuest(client.Eclient, evntID, leavingUser)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	return

}
