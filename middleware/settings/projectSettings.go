package settings

import (
	"log"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Project ...
func Project(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projURL := r.URL.Path[17:]
	// projURL, err := get.ProjectByID(client.Eclient, r.FormValue("projectID"))
	project, err := uses.AggregateProjectData(client.Eclient, projURL, test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		log.Println(err)
	}

	var isAdmin = false
	for _, member := range project.ProjectData.Members {
		if member.MemberID == test1.(string) && member.Role <= 0 {
			isAdmin = true
			break
		}
	}
	if isAdmin {
		cs := client.ClientSide{Project: project}
		client.RenderSidebar(w, r, "template2-nil")
		client.RenderSidebar(w, r, "leftnav-nil")
		client.RenderTemplate(w, r, "project_settings_F", cs)

	} else {
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

}
