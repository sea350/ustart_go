package project

import (
	"fmt"
	"net/http"

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

	projectID := r.FormValue("UNKOWN")
	newContent := []rune(r.FormValue("content"))
	err := uses.ProjectCreatesEntry(client.Eclient, projectID, docID.(string), newContent)
	if err != nil {
		fmt.Println("err: middleware/project/makeentries line 26")
		fmt.Println(err)
	}

	fmt.Fprintln(w, "complete")
}
