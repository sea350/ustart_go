package settings

import (
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	uses "github.com/sea350/ustart_go/uses"
)

//LeaveProject ... lets a user leave a project
//If Rol
func LeaveProject(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	leavingUser := r.FormValue("leaverID")
	projID := r.FormValue("projectID")
	newCreator := r.FormValue("newCreator")

	proj, err := get.ProjectByID(eclient, projID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	var canLeave = false
	if leavingUser == test1.(string) {
		//if the current active user wants to leave, they can
		canLeave = true
	} else {
		for _, mem := range proj.Members {
			if mem.MemberID == test1.(string) && mem.Role == 0 {
				//if the current acessing user is creator, they can do whatever they want
				canLeave = true
				break
			}
		}
	}
	if !canLeave {
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	if newCreator == `` {
		err = post.DeleteMember(client.Eclient, projID, leavingUser)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	} else {
		err = uses.NewProjectLeader(client.Eclient, projID, leavingUser, newCreator)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		err = post.DeleteMember(client.Eclient, projID, leavingUser)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return

}
