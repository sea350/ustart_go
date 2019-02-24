package project

import (
	"html"

	"net/http"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//AddQuickLink ...
func AddQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID := r.FormValue("projectID")
	if ID == `` {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Crucial data was not passed in, now exiting")
		return
	}

	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	var exists bool
	var member types.Member
	for _, mem := range proj.Members {
		if mem.MemberID == session.Values["DocID"].(string) {
			exists = true
			member = mem
			break
		}
	}

	if !exists {
		return
	}

	// hasPermission := uses.HasPrivilege("links", proj.PrivilegeProfiles, member)

	// if !hasPermission {
	// 	return
	// }

	p := bluemonday.UGCPolicy()
	cleanProjHTML := p.Sanitize(r.FormValue("projectLink"))
	if len(cleanProjHTML) == 0 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Link cannot be blank")
	}
	isValid := uses.ValidLink(cleanProjHTML)
	if !isValid {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Invalid link provided")
		return
	}
	cleanTitle := p.Sanitize(r.FormValue("projectLinkDesc"))
	if len(cleanTitle) == 0 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Title cannot be blank")
	}
	proj.QuickLinks = append(proj.QuickLinks, types.Link{Name: html.EscapeString(cleanTitle), URL: html.EscapeString(cleanProjHTML)})

	err = post.UpdateProject(client.Eclient, ID, "QuickLinks", proj.QuickLinks)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

}
