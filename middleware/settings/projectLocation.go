package settings

import (
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectLocation ...
//For Projects Location
func ProjectLocation(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()
	var vis bool

	p := bluemonday.UGCPolicy()
	visibool := r.FormValue("locVis")
	cleanVis := p.Sanitize(visibool)
	if cleanVis == "true" {
		vis = true
	} else {
		vis = false
	}
	cleanCountry := p.Sanitize(r.FormValue("country"))
	country := cleanCountry
	cleanState := p.Sanitize(r.FormValue("state"))
	state := cleanState
	cleanCity := p.Sanitize(r.FormValue("city"))
	city := cleanCity
	cleanZip := p.Sanitize(r.FormValue("zip"))
	zip := cleanZip

	projID := r.FormValue("projectID")

	proj, err := get.ProjectByID(client.Eclient, projID)
	//TODO: DocID

	err = uses.ChangeProjectLocation(client.Eclient, projID, country, state, city, zip, vis)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	http.Redirect(w, r, "/ProjectSettings/"+proj.URLName, http.StatusFound)
	return

}
