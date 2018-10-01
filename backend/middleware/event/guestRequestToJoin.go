package event

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/backend/get/event"
	client "github.com/sea350/ustart_go/backend/middleware/client"
	userPost "github.com/sea350/ustart_go/backend/post/user"
)

//GuestRequestToJoin ... PUBLIC no need for request received
func GuestRequestToJoin(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	id := r.FormValue("eventID") //event docID
	evnt, err := get.EventByID(client.Eclient, id)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	for _, guestInfo := range evnt.Guests {
		if guestInfo.GuestID == test1.(string) {
			http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
			return
		}
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("event ID", id)

	err = userPost.AppendSentEventReq(client.Eclient, test1.(string), id)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)

}
