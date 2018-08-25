package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeNameAndDescription ...
//For Projects
func ChangeNameAndDescription(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	projName := r.FormValue("pname")
	projDesc := []rune(r.FormValue("inputDesc"))
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
