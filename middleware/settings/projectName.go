package settings

import (
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeNameAndDescription ...
//For Projects
func ChangeNameAndDescription(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := bluemonday.UGCPolicy()
	r.ParseForm()
	projName := p.Sanitize(r.FormValue("pname"))
	if len(projName) < 1 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Project name cannot be blank!")
		return
	}
	projDesc := []rune(p.Sanitize(r.FormValue("inputDesc")))

	proj, err := get.ProjectByID(client.Eclient, r.FormValue("projectID"))
	//TODO: DocID
	err = uses.ProjectNameAndDescription(client.Eclient, r.FormValue("projectID"), projName, projDesc)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	//TODO: Add in right URL
	http.Redirect(w, r, "/ProjectSettings/"+proj.URLName, http.StatusFound)
	return

}
