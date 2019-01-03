package event

import (
	"encoding/json"
	"fmt"
	
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	getEvent "github.com/sea350/ustart_go/get/event"
	"github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/uses"
)

//MakeEventEntry ... used to make an original entry for events made for ajax
func MakeEventEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		//No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := bluemonday.UGCPolicy()

	eventID := p.Sanitize(r.FormValue("docID"))
	if eventID == "" {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Critical data not passed in")
	}
	newContent := []rune(p.Sanitize(r.FormValue("text")))

	event, err := getEvent.EventByID(client.Eclient, eventID)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		return
	}

	var isMem bool
	for _, mem := range event.Members {
		if mem.MemberID == docID.(string) {
			isMem = true
			break
		}
	}

	if !isMem {
		return
	}

	newID, err := uses.EventCreatesEntry(client.Eclient, eventID, docID.(string), newContent)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		return
	}

	jEntry, err := uses.ConvertEntryToJournalEntry(client.Eclient, newID, docID.(string), true)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	data, err := json.Marshal(jEntry)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	fmt.Fprintln(w, string(data))
}
