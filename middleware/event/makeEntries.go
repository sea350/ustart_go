package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/uses"
)

//MakeEventEntry ... used to make an original entry for events made for ajax
func MakeEventEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		//No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	eventID := r.FormValue("docID")
	newContent := []rune(r.FormValue("text"))
	newID, err := uses.EventCreatesEntry(client.Eclient, eventID, docID.(string), newContent)
	if err != nil {
		fmt.Println("err: middleware/event/makeentries line 26")
		fmt.Println(err)
	}

	jEntry, err := uses.ConvertEntryToJournalEntry(client.Eclient, newID, true)
	if err != nil {
		fmt.Println("err: middleware/event/makeentries line 32")
		fmt.Println(err)
	}

	data, err := json.Marshal(jEntry)
	if err != nil {
		fmt.Println("err: middleware/event/makeentries line 38")
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
