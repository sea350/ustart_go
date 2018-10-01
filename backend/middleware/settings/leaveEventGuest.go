package settings

import (
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"

	get "github.com/sea350/ustart_go/backend/get/event"
	client "github.com/sea350/ustart_go/backend/middleware/client"
	post "github.com/sea350/ustart_go/backend/post/event"
)

//LeaveEventGuest ... lets a user leave a event
//If Rol
func LeaveEventGuest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()
	leavingUser := p.Sanitize(r.FormValue("leaverID"))
	evntID := p.Sanitize(r.FormValue("eventID"))

	evnt, err := get.EventByID(client.Eclient, evntID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	return

}
