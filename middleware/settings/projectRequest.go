package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	uses "github.com/sea350/ustart_go/uses"
	types "github.com/sea350/ustart_go/types"

)

//LeaveProject ... lets a user leave a project
//If Rol
func ProjectRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	r.ParseForm()
	requestUser := r.FormValue("username")
	isAccepted := true
	projID := r.FormValue("projectID")
	proj, err := get.ProjectByID(eclient, projID)
	if err != nil {
		fmt.Println(err)
	}

	err = uses.RemoveRequest(eclient, projID, requestUser)

	if err != nil {
		fmt.Println(err)
	}

	if isAccepted {
		var newInfo types.ProjectInfo{}
		newInfo.ProjectID = projID
		newInfo.Visible = true
		err = uses.AcceptRequest(eclient, newInfo, "" )
	}

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)

}
