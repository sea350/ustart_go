package project

import (
	"encoding/json"
	"html"

	get "github.com/sea350/ustart_go/get/project"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"

	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
)

//UpdateTags ...
func UpdateTags(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()

	ID := r.FormValue("skillArray")
	theID := r.FormValue("projectWidget")

	proj, err := get.ProjectByID(client.Eclient, theID)
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

	hasPermission := uses.HasPrivilege("tags", proj.PrivilegeProfiles, member)

	if !hasPermission {
		return
	}

	var ts []string
	err := json.Unmarshal([]byte(ID), &ts)

	if err != nil {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Could not unmarshal")
		return
	}

	var updatedTags []string
	for t := range ts {
		ts[t] = p.Sanitize(ts[t])
		ts[t] = html.EscapeString(ts[t])
		allowed := uses.TagAllowed(client.Eclient, ts[t])
		if allowed {
			updatedTags = append(updatedTags, ts[t])
		}

	}

	err = post.UpdateProject(client.Eclient, theID, "Tags", updatedTags)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
