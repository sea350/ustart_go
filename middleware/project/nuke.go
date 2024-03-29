package project

import (
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
)

//Nuke ... kills all project dpendancies and sets project to invisible (acts like delete project)
//designed for ajax
func Nuke(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	defer http.Redirect(w, r, "/", http.StatusFound)

	projID := r.FormValue("projectID")
	if projID == `` {

		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "No project ID passed in")
		return
	}

	proj, err := get.ProjectByID(client.Eclient, projID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	var permission bool
	for _, mem := range proj.Members {
		if mem.MemberID == docID.(string) {
			if mem.Role == 0 {
				permission = true
			}
			break
		}
	}

	if !permission {

		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "User ID " + docID.(string) + "does not have authorization to delete project " + projID)
		return
	}
	err = post.InvisProject(client.Eclient, projID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
