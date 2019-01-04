package event

import (
	"encoding/json"
	"fmt"
	
	"net/http"

	"github.com/sea350/ustart_go/uses"

	get "github.com/sea350/ustart_go/get/event"
	"github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
)

//LoadGuestJoinRequests ...
func LoadGuestJoinRequests(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID := r.FormValue("eventID") //eventID
	if ID == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Event ID not passed in")
		return
	}

	var heads []types.FloatingHead

	evnt, err := get.EventByID(client.Eclient, ID)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	for index, userID := range evnt.GuestReqReceived {
		//user
		if userID == 1 {
			head, err := uses.ConvertUserToFloatingHead(client.Eclient, index)
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}
			heads = append(heads, head)
		}
		//project
		if userID == 2 {
			head, err := uses.ConvertProjectToFloatingHead(client.Eclient, index)
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}
			heads = append(heads, head)
		}

	}

	data, err := json.Marshal(heads)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	fmt.Fprintln(w, string(data))

}
