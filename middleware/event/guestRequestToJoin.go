package event

import (
	"fmt"
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	evntPost "github.com/sea350/ustart_go/post/event"
	userPost "github.com/sea350/ustart_go/post/user"
)

//GuestRequestToJoin ...
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
	for _, receivedReq := range evnt.GuestReqReceived {
		if receivedReq == test1.(string) {
			http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
			return
		}
	}
	fmt.Println("event ID", id)
	fmt.Println("user ID", test1.(string))


	"""
	if _, exists := evnt.GuestReqReceived[test1.(string)]; exists {
		fmt.Println("GuestReqReceived working?")
		http.Redirect(w, r, "\/Event/\"+evnt.URLName, http.StatusFound)
		return
	}
	"""

	err = userPost.AppendSentEventReq(client.Eclient, test1.(string), id)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	fmt.Println("AppendSentEventReq working?", err)
	err = evntPost.AppendGuestReqReceived(client.Eclient, id, test1.(string), evnt.GuestReqReceived[test1.(string)])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	fmt.Println("AppendGuestReqReceived working?", err)

	http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
	return
}
