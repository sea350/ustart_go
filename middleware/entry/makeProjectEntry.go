package entry

import (
	"encoding/json"
	"fmt"
	"html"

	"github.com/sea350/ustart_go/antispam"

	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	postEntry "github.com/sea350/ustart_go/post/entry"
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
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	projectID := r.FormValue("docID")
	if projectID == `` {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Critical data not passed in")
		return
	}
	newContent := html.EscapeString(r.FormValue("text"))
	proj, member, err := get.ProjAndMember(client.Eclient, projectID, docID.(string))

	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	if !antispam.AntiJournalSpam(docID.(string)) {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | This user is attempting to spam on project " + projectID)
		return
	}

	if uses.HasPrivilege("post", proj.PrivilegeProfiles, member) {
		var entry types.Entry
		entry.ProjectOriginalEntry(docID.(string), projectID, newContent)

		entryID, err := postEntry.IndexEntry(client.Eclient, entry)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			return
		}

		// err = post.AppendEntryID(client.Eclient, projectID, entryID)
		// if err != nil {

		// 	client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		// }

		jEntry, err := uses.ConvertEntryToJournalEntry(client.Eclient, entryID, docID.(string), true)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		data, err := json.Marshal(jEntry)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		fmt.Fprintln(w, string(data))
	} else {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | err: This user does not have permission to perform this action")
	}
}
