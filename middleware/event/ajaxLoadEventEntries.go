package event

import (
	"encoding/json"
	"fmt"
	"io"
	
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	client "github.com/sea350/ustart_go/middleware/client"
	scrollpkg "github.com/sea350/ustart_go/properloading"
)

//AjaxLoadEventEntries ... pulls all entries for a given event and fprints it back as a json array
func AjaxLoadEventEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()

	wallID := p.Sanitize(r.FormValue("userID"))
	/*
		evnt, err := get.EventByID(client.Eclient, wallID)
		if err != nil {
			
	
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}


			entries, err := uses.LoadEntries(client.Eclient, evnt.EntryIDs)
			if err != nil {
				
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			}
	*/

	res, entries, total, err := scrollpkg.ScrollPageEvents(client.Eclient, []string{wallID}, docID.(string), "")
	if err != nil && err != io.EOF {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	results := make(map[string]interface{})
	results["JournalEntries"] = entries
	results["ScrollID"] = res
	results["TotalHits"] = total

	data, err := json.Marshal(results)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	fmt.Fprintln(w, string(data))
}
