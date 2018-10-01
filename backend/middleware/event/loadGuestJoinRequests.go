package event

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sea350/ustart_go/backend/uses"

	get "github.com/sea350/ustart_go/backend/get/event"
	"github.com/sea350/ustart_go/backend/middleware/client"
	types "github.com/sea350/ustart_go/backend/types"
)

//LoadGuestJoinRequests ...
func LoadGuestJoinRequests(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	ID := r.FormValue("eventID") //eventID

	var heads []types.FloatingHead

	evnt, err := get.EventByID(client.Eclient, ID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	for index, userID := range evnt.GuestReqReceived {
		if userID == 1 {
			head, err := uses.ConvertUserToFloatingHead(client.Eclient, index)
			if err != nil {
				fmt.Println(err)
				fmt.Println(fmt.Sprintf("err: middleware/event/loadjoinrequest, Line 35, index %s", index))
			}
			heads = append(heads, head)
		}
		if userID == 2 {
			head, err := uses.ConvertProjectToFloatingHead(client.Eclient, index)
			if err != nil {
				fmt.Println(err)
				fmt.Println(fmt.Sprintf("err: middleware/event/loadprojjoinrequest, Line 40, index %s", index))
			}
			heads = append(heads, head)
		}

	}

	data, err := json.Marshal(heads)
	if err != nil {
		fmt.Println("err: middleware/project/loadjoinrequest, Line 45")
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(data))

}
