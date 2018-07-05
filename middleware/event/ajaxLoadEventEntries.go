package event

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Println("Error: middleware/event/AjaxLoadEventEntries line 23")
		log.Println(err)
	}

	entries, err := uses.LoadEntries(client.Eclient, evnt.EntryIDs)
	if err != nil {
		log.Println("Error: middleware/event/AjaxLoadEventEntries line 29")
		log.Println(err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		log.Println("Error: middleware/event/AjaxLoadEventEntries line 35")
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
