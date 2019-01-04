package project

import (
	"fmt"
	
	"net/http"
	

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//RejectJoinRequest ...
func RejectJoinRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	projID := r.FormValue("projectID")
	newMemberID := r.FormValue("userID")

	newNumRequests, err := uses.RemoveRequest(client.Eclient, projID, newMemberID)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	fmt.Fprintln(w, newNumRequests)
}
