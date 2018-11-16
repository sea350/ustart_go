package entry

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	client "github.com/sea350/ustart_go/middleware/client"
	postEntry "github.com/sea350/ustart_go/post/entry"
	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//MakeUserEntry ... makes user original post
//designed for ajax
func MakeUserEntry(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()
	p := bluemonday.UGCPolicy()

	text := p.Sanitize(r.FormValue("text"))

	var entry types.Entry
	entry.UserOriginalEntry(docID.(string), text)

	entryID, err := postEntry.IndexEntry(client.Eclient, entry)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	jEntry, err := uses.ConvertEntryToJournalEntry(client.Eclient, entryID, docID.(string), true)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	data, err := json.Marshal(jEntry)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(data))
}
