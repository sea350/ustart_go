package event

import (
	"fmt"
	"log"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	evntPost "github.com/sea350/ustart_go/post/event"
	userPost "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//AcceptGuestJoinRequest ...
func AcceptGuestJoinRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	evntID := r.FormValue("eventID")
	newGuestID := r.FormValue("userID")

	newNumRequests, err := uses.RemoveGuestRequest(client.Eclient, evntID, newGuestID)
	if err != nil {
		log.Println("Error: middleware/event/acceptGuestJoinRequest line 27")
		log.Println(err)
	}

	err = userPost.AppendEvent(client.Eclient, newGuestID, types.EventInfo{EventID: evntID, Visible: true})
	if err != nil {
		log.Println("Error: middleware/event/acceptGuestJoinRequest line 33")
		log.Println(err)
	}

	var newGuest types.EventGuests
	newGuest.GuestID = newGuestID
	newGuest.Visible = true

	err = evntPost.AppendGuest(client.Eclient, evntID, newGuest)
	if err != nil {
		log.Println("Error: middleware/event/acceptGuestJoinRequest line 43")
		log.Println(err)
	}

	fmt.Fprintln(w, newNumRequests)

}
