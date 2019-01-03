package settings

import (
	"net/http"

	"github.com/microcosm-cc/bluemonday"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
	uses "github.com/sea350/ustart_go/uses"
)

//LeaveEvent ... lets a user leave a event
//If Rol
func LeaveEvent(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()
	leavingUser := p.Sanitize(r.FormValue("leaverID"))
	if len(leavingUser) < 1 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "This field cannot be left blank!")
		return
	}
	eventID := p.Sanitize(r.FormValue("eventID"))
	// if len(eventID) < 1{
	// 			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"This field cannot be left blank!")
	// 	return
	// }
	newCreator := p.Sanitize(r.FormValue("newCreator"))
	if len(newCreator) < 1 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "This field cannot be left blank!")
		return
	}

	event, err := get.EventByID(client.Eclient, eventID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	var canLeave = false
	if leavingUser == test1.(string) {
		//if the current active user wants to leave, they can
		canLeave = true
	} else {
		for _, mem := range event.Members {
			if mem.MemberID == test1.(string) && mem.Role == 0 {
				//if the current acessing user is creator, they can do whatever they want
				canLeave = true
				break
			}
		}
	}
	if !canLeave {
		http.Redirect(w, r, "/Event/"+event.URLName, http.StatusFound)
		return
	}

	if newCreator == `` {
		err = post.DeleteMember(client.Eclient, eventID, leavingUser)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}
	} else {
		err = uses.NewEventLeader(client.Eclient, eventID, leavingUser, newCreator)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}
		err = post.DeleteMember(client.Eclient, eventID, leavingUser)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}
	}

	http.Redirect(w, r, "/Event/"+event.URLName, http.StatusFound)
	return

}
