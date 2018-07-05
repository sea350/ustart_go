package event

import (
	"fmt"
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

	ID := r.FormValue("eventID") //event docID
	fmt.Println(ID)
	fmt.Println("debug text requesttojoin line 23")

	evnt, err := get.EventByID(client.Eclient, ID)
	if err != nil {
		fmt.Println("err middleware/event/guestrequesttojoin line25")
		fmt.Println(err)
	}

	for _, guestInfo := range evnt.Guests {
		if guestInfo.GuestID == test1.(string) {
			http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
			return
		}
	}
	for _, receivedReq := range evnt.GuestReqReceived {
		if receivedReq == test1.(string) {
			http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
			return
		}
	}
	err = userPost.AppendSentEventReq(client.Eclient, test1.(string), ID)
	if err != nil {
		fmt.Println("err middleware/event/guestrequesttojoin line42")
		fmt.Println(err)
	}
	err = evntPost.AppendGuestReqReceived(client.Eclient, ID, test1.(string))
	if err != nil {
		fmt.Println("err middleware/event/guestrequesttojoin line47")
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	return
}
