package project

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/uses"
)

//MakeEntry ... used to make an original entry for projects
//made for ajax
func MakeEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projectID := r.FormValue("docID")
	newContent := []rune(r.FormValue("text"))
	proj, member, err := get.ProjAndMember(client.Eclient, projectID, docID.(string))

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	if uses.HasPrivilege("post", proj.PrivilegeProfiles, member) {
		newID, err := uses.ProjectCreatesEntry(client.Eclient, projectID, docID.(string), newContent)

		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}

		jEntry, err := uses.ConvertEntryToJournalEntry(client.Eclient, newID, true)
		if err != nil {
			fmt.Println("err: middleware/project/makeentries line 32")
			fmt.Println(err)
		}
		data, err := json.Marshal(jEntry)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		fmt.Fprintln(w, string(data))
	}
}
