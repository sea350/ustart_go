package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		fmt.Println(test1)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := bluemonday.UGCPolicy()
	r.ParseForm()
	projName := p.Sanitize(r.FormValue("pname"))
	if len(projName) < 1 {
		log.Println("Project name cannot be blank!")
		return
	}
	projDesc := []rune(p.Sanitize(r.FormValue("inputDesc")))
	//   fmt.Println(blob)
	fmt.Println(projName, projName)

	//fmt.Println(reflect.TypeOf(blob))
	proj, err := get.ProjectByID(client.Eclient, r.FormValue("projectID"))
	//TODO: DocID
	err = uses.ProjectNameAndDescription(client.Eclient, r.FormValue("projectID"), projName, projDesc)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	//TODO: Add in right URL
	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return

}
