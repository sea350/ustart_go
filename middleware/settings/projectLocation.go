package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectLocation ...
//For Projects Location
func ProjectLocation(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["projectID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	//   fmt.Println(blob)

	projID := r.FormValue("projectID")
	proj, err := get.ProjectByID(eclient, projID)
	//fmt.Println(reflect.TypeOf(blob))
	//TODO: DocID
	err = uses.ChangeProjectLocation(eclient, projID, country, state, city, zip)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	//TODO: Add in right URL
	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return

}
