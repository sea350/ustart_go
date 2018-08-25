package event

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
	types "github.com/sea350/ustart_go/types"
)

//NewGuest ...processes a request to add a guest, will check for all contingencies... eventually
//Designed for ajax
func NewGuest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	sessionID, _ := session.Values["DocID"]
	if sessionID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	eventID := r.FormValue("eventID")

	event, err := get.EventByID(client.Eclient, eventID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)

	}

	for i := range event.Guests {
		if event.Guests[i].GuestID == sessionID.(string) {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("User attempting to join is already a guest")
			http.Redirect(w, r, "/Event/"+event.URLName, http.StatusFound)
			return
		}
	}
	err = post.AppendGuest(client.Eclient, eventID, types.EventGuests{GuestID: sessionID.(string), Status: 1})
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	http.Redirect(w, r, "/Event/"+event.URLName, http.StatusFound)
	return

}
