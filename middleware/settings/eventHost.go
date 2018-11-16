package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/event"
	getproj "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//EventHost ...
func EventHost(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["eventID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()
	projectID := r.FormValue("projectID")
	evntID := r.FormValue("eventID")
	evnt, err := get.EventByID(client.Eclient, evntID)
	if err != nil {
		panic(err)
	}
	_, err = getproj.ProjectByID(client.Eclient, projectID)
	if err == nil {
		err = uses.ChangeEventHost(client.Eclient, evntID, projectID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	}
	//TODO: Add in right URL
	http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	return

}
