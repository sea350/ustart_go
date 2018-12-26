package event

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}


			entries, err := uses.LoadEntries(client.Eclient, evnt.EntryIDs)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
			}
	*/

	res, entries, total, err := scrollpkg.ScrollPageEvents(client.Eclient, []string{wallID}, docID.(string), "")
	if err != nil && err != io.EOF {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	results := make(map[string]interface{})
	results["JournalEntries"] = entries
	results["ScrollID"] = res
	results["TotalHits"] = total

	data, err := json.Marshal(results)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
