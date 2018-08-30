package entry

import (
	"fmt"
	"log"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	postEntry "github.com/sea350/ustart_go/post/entry"
	"github.com/sea350/ustart_go/types"
)

//ShareEntry ... Creates a new shared entry for user
func ShareEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	originalPost := r.FormValue("postid")
	newContent := r.FormValue("content")

	var entry types.Entry
	entry.UserShareEntry(docID.(string), originalPost, newContent)

	replyID, err := postEntry.IndexEntry(client.Eclient, entry)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	err = postEntry.AppendShareID(client.Eclient, originalPost, replyID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, originalPost)
}
