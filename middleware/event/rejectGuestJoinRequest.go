package event

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//RejectEventGuestJoinRequest ...
func RejectEventGuestJoinRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	evntID := r.FormValue("eventID")
	if evntID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Critical information not passed in")
		return
	}
	newMemberID := r.FormValue("userID")
	if newMemberID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Critical information not passed in")
		return
	}
	classification, err := strconv.Atoi(r.FormValue("classification"))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	newNumRequests, err := uses.RemoveGuestRequest(client.Eclient, evntID, newMemberID, classification)
	if err != nil {
		fmt.Println("err middleware/event/rejectguestjoinrequest line 27")
		fmt.Println(err)
	}

	fmt.Fprintln(w, newNumRequests)
}
