package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//AjaxLoadEventEntries ... pulls all entries for a given event and fprints it back as a json array
func AjaxLoadEventEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	wallID := r.FormValue("userID")
	evnt, err := get.EventByID(client.Eclient, wallID)
	if err != nil {
		fmt.Println("err middleware/event/AjaxLoadEventEntries line 25")
		fmt.Println(err)
	}

	entries, err := uses.LoadEntries(client.Eclient, evnt.EntryIDs)
	if err != nil {
		fmt.Println("err middleware/event/AjaxLoadEventEntries line 30")
		fmt.Println(err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("err middleware/event/AjaxLoadEventEntries line 37")
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
