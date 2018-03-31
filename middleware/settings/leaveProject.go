package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//LeaveProject ... lets a user leave a project
//If Rol
func LeaveProject(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	leavingUser := r.FormValue("UNKOWN")
	projID := r.FormValue("UNKNOWN")
	newLeader := r.FormValue("UNKNOWN")

	proj, err := get.ProjectByID(eclient, projID)
	if err != nil {
		fmt.Println("err middleware/settings/leaveproject line 28")
		fmt.Println(err)
	}

	isLeader, _ := uses.IsLeader(eclient, projID, test1.(string))
	if isLeader {
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	if newLeader != `` {
		err = uses.RemoveMember(client.Eclient, projID, leavingUser)
		fmt.Println("err middleware/settings/leaveproject line 34")
		fmt.Println(err)
	} else {
		err = uses.NewProjectLeader(client.Eclient, projID, leavingUser, newLeader)
		fmt.Println("err middleware/settings/leaveproject line 38")
		fmt.Println(err)
		err = uses.RemoveMember(client.Eclient, projID, leavingUser)
		fmt.Println("err middleware/settings/leaveproject line 41")
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)

}
