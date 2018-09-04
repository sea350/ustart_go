package settings

import (
	"fmt"
	"log"
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
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	var vis bool
	log.Println("SETTING UGCPOLICY")
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
	//   fmt.Println(blob)

	projID := r.FormValue("projectID")
	fmt.Println("projiD", projID)
	fmt.Println("city", city)
	fmt.Println("city", zip)
	log.Println("GETTING PROJECT BY ID")
	proj, err := get.ProjectByID(client.Eclient, projID)
	//fmt.Println(reflect.TypeOf(blob))
	//TODO: DocID
	log.Println("CHANGING PROJECT LOCATION")
	err = uses.ChangeProjectLocation(client.Eclient, projID, country, state, city, zip, vis)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		log.Println(err)
	}

	fmt.Println("/Projects/", proj.URLName)
	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return

}
