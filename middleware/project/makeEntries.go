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
	err := uses.ProjectCreatesEntry(client.Eclient, projectID, newContent)
	if err != nil {
		fmt.Println("err: middleware/project/makeentries line 26")
		fmt.Println(err)
	}

	fmt.Fprintln(w, "complete")
}

//ReplyEntry ... makes a reply entry for projects
//made for ajax
func ReplyEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projectID := r.FormValue("UNKOWN")
	originalPost := r.FormValue("postid")
	newContent := []rune(r.FormValue("content"))
	err := uses.ProjectCreatesReply(client.Eclient, projectID, originalPost, newContent)
	if err != nil {
		fmt.Println("err: middleware/project/makeentries line 49")
		fmt.Println(err)
	}

	fmt.Fprintln(w, "complete")
}

//ShareEntry ... makes a share entry for projects
//made for ajax
func ShareEntry(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projectID := r.FormValue("UNKOWN")
	originalPost := r.FormValue("postid")
	newContent := []rune(r.FormValue("content"))
	err := uses.ProjectCreatesShare(client.Eclient, projectID, originalPost, newContent)
	if err != nil {
		fmt.Println("err: middleware/project/makeentries line 72")
		fmt.Println(err)
	}

	fmt.Fprintln(w, "complete")
}
