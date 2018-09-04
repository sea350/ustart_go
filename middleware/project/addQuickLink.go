package project

import (
	"html"
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/types"
)

//AddQuickLink ...
func AddQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	ID := r.FormValue("projectID")

	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	p := bluemonday.UGCPolicy()
	cleanProjHTML := p.Sanitize(r.FormValue("projectLink"))
	proj.QuickLinks = append(proj.QuickLinks, types.Link{Name: html.EscapeString(r.FormValue("projectLinkDesc")), URL: cleanProjHTML})

	err = post.UpdateProject(client.Eclient, ID, "QuickLinks", proj.QuickLinks)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

}
