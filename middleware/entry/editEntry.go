package entry

import (
	"encoding/json"
	"fmt"
	
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/entry"
	uses "github.com/sea350/ustart_go/uses"
)

//EditEntry ... edits entry NEEDS SECURITY REVISIT
//designed for ajax
func EditEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := bluemonday.UGCPolicy()

	postID := r.FormValue("postid")
	newContent := p.Sanitize(r.FormValue("content"))

	err := post.UpdateEditEntry(client.Eclient, postID, "Content", []rune(newContent))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	jEntry, err := uses.ConvertEntryToJournalEntry(client.Eclient, postID, docID.(string), true)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	data, err := json.Marshal(jEntry)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	fmt.Fprintln(w, string(data))

}
