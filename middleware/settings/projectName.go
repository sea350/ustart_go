package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeNameAndDescription ...
//For Projects
func ChangeNameAndDescription(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	r.ParseForm()
	projName := r.FormValue("pname")
	projDesc := []rune(r.FormValue("inputDesc"))
	//   fmt.Println(blob)
	fmt.Println(projName, projName)

	//fmt.Println(reflect.TypeOf(blob))
	proj, err := get.ProjectByID(eclient, r.FormValue("projectID"))
	//TODO: DocID
	err = uses.ProjectNameAndDescription(eclient, r.FormValue("projectID"), projName, projDesc)

	if err != nil {
		fmt.Println(err)
	}
	//TODO: Add in right URL
	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)

}
