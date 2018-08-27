package entry

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	postEntry "github.com/sea350/ustart_go/post/entry"
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//MakeProjectEntry ... used to make an original entry for projects
//made for ajax
func MakeProjectEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projectID := r.FormValue("docID")
	newContent := html.EscapeString(r.FormValue("text"))
	proj, member, err := get.ProjAndMember(client.Eclient, projectID, docID.(string))

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	if uses.HasPrivilege("post", proj.PrivilegeProfiles, member) {
		var entry types.Entry
		entry.ProjectOriginalEntry(docID.(string), projectID, newContent)

		entryID, err := postEntry.IndexEntry(client.Eclient, entry)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}

		err = post.AppendEntryID(client.Eclient, projectID, entryID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
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
}
