package project

import (
	"encoding/json"

	"net/http"

	"github.com/microcosm-cc/bluemonday"

	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//UpdateSkills ...
func UpdateSkills(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID := r.FormValue("projectWidget")
	skills := r.FormValue("skillArray")

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

	hasPermission := uses.HasPrivilege("tags", proj.PrivilegeProfiles, member)

	if !hasPermission {
		return
	}
	p := bluemonday.UGCPolicy()

	var ss []string
	err = json.Unmarshal([]byte(skills), &ss)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	for s := range ss {
		ss[s] = p.Sanitize(ss[s])
	}
	err = post.UpdateProject(client.Eclient, ID, "ListNeeded", ss)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
