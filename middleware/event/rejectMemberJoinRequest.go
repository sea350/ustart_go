package event

import (
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//RejectEventMemberJoinRequest ...
func RejectEventMemberJoinRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	evntID := r.FormValue("eventID")
	newMemberID := r.FormValue("userID")

	newNumRequests, err := uses.RemoveEventRequest(client.Eclient, evntID, newMemberID)
	if err != nil {
		fmt.Println("err middleware/event/rejectmemberjoinrequest line 27")
		fmt.Println(err)
	}

	fmt.Fprintln(w, newNumRequests)
}
