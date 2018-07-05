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
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	ID := r.FormValue("eventID") //eventID
	fmt.Println("debug text middlware/event/loadjoinrequests line 24")
	fmt.Println(ID)
	var heads []types.FloatingHead

	evnt, err := get.EventByID(client.Eclient, ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("err: middleware/event/loadjoinrequest Line 26")
	}

	for index, userID := range evnt.GuestReqReceived {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, userID)
		if err != nil {
			fmt.Println(err)
			fmt.Println(fmt.Sprintf("err: middleware/event/loadjoinrequest, Line 35, index %d", index))
		}
		heads = append(heads, head)
	}

	data, err := json.Marshal(heads)
	if err != nil {
		fmt.Println("err: middleware/project/loadjoinrequest, Line 45")
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(data))

}
