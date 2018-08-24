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
	id := r.FormValue("docID") //event docID
	fmt.Println("ID", id)
	evnt, err := get.EventByID(client.Eclient, id)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	fmt.Println("test1 value", test1.(string))
	fmt.Println("event url name", evnt.URLName)
	for _, guestInfo := range evnt.Guests {
		fmt.Println("guestinfo. guestid", guestInfo.GuestID)
		if guestInfo.GuestID == test1.(string) {
			http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
			return
		}
	}

	if _, exists := evnt.GuestReqReceived[test1.(string)]; exists {
		http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
		return
	}

	err = userPost.AppendSentEventReq(client.Eclient, test1.(string), ID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	err = evntPost.AppendGuestReqReceived(client.Eclient, ID, test1.(string), evnt.GuestReqReceived[test1.(string)])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	http.Redirect(w, r, "/Event/"+evnt.URLName, http.StatusFound)
	return
}
