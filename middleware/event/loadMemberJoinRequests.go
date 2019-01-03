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

//LoadMemberJoinRequests ...
func LoadMemberJoinRequests(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID := r.FormValue("eventID") //eventID

	var heads []types.FloatingHead

	evnt, err := get.EventByID(client.Eclient, ID)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	for index, userID := range evnt.MemberReqReceived {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, userID)
		if err != nil {
			
					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+fmt.Sprintf("error loading index %d", index))
		}
		heads = append(heads, head)
	}

	data, err := json.Marshal(heads)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	fmt.Fprintln(w, string(data))

}
