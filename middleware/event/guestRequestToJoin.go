package event

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
)

//GuestRequestToJoin ... PUBLIC no need for request received
func GuestRequestToJoin(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	id := r.FormValue("eventID") //event docID
	if id == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Event ID not passed in")
		return
	}

	evnt, err := get.EventByID(client.Eclient, id)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("Debug check 1/2")

	for _, guestInfo := range evnt.Guests {
		if guestInfo.GuestID == test1.(string) {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("user is already a guest")
			http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
			return
		}
	}
	for _, receivedReq := range evnt.MemberReqReceived {
		if receivedReq == test1.(string) {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("user has already sent a request")
			http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
			return
		}
	}

	err = post.AppendGuestReqReceived(client.Eclient, id, test1.(string), 1)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("Debug check 2/2")

	http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)

}
