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
	if evntID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("WARNING: event ID not received")
		return
	}
	newGuestID := r.FormValue("userID")
	if newGuestID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("WARNING: new member ID not received")
		return
	}
	//classification, err := strconv.Atoi(r.FormValue("classification")) GUEST ARE classification 1 right??

	newNumRequests, err := uses.RemoveGuestRequest(client.Eclient, evntID, newGuestID, 1)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	err = userPost.AppendEvent(client.Eclient, newGuestID, types.EventInfo{EventID: evntID, Visible: true})
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	var newGuest types.EventGuests
	newGuest.Status = 0
	newGuest.GuestID = newGuestID
	newGuest.Classification = 1

	err = evntPost.AppendGuest(client.Eclient, evntID, newGuest)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, newNumRequests)

}
