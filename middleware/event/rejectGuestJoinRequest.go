package event

import (
	"fmt"
	
	"net/http"

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
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Critical information not passed in")
		return
	}
	newMemberID := r.FormValue("userID")
	if newMemberID == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Critical information not passed in")
		return
	}
	// classification, err := strconv.Atoi(r.FormValue("classification"))
	// if err != nil {
	// 	
	// 	client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	// 	return
	// }

	newNumRequests, err := uses.RemoveGuestRequest(client.Eclient, evntID, newMemberID, 1)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	fmt.Fprintln(w, newNumRequests)
}
