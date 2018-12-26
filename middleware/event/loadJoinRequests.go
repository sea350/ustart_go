package event

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	get "github.com/sea350/ustart_go/get/event"
	"github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
)

//LoadJoinRequests ...
func LoadJoinRequests(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID := r.FormValue("eventID") //eventID
	if ID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Event ID not passed in")
		return
	}

	var heads []types.FloatingHead

	evnt, err := get.EventByID(client.Eclient, ID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	for index, userID := range evnt.GuestReqReceived {
		//user
		if userID == 1 {
			head, err := uses.ConvertUserToFloatingHead(client.Eclient, index)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
			heads = append(heads, head)
		}
		//project
		if userID == 2 {
			head, err := uses.ConvertProjectToFloatingHead(client.Eclient, index)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
			heads = append(heads, head)
		}

	}

	var heads2 []types.FloatingHead

	for index, userID := range evnt.MemberReqReceived {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, userID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(fmt.Sprintf("error : problem loading index %d", index))
		}
		heads2 = append(heads2, head)
	}

	marshall := make(map[string][]types.FloatingHead)
	marshall["GuestRequests"] = heads
	marshall["MemberRequests"] = heads2

	data, err := json.Marshal(marshall)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))

}
