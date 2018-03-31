package settings

import (
	"fmt"
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
	}

	projID := r.URL.Path[17:]
	fmt.Println(projID)
	fmt.Println("debug text projectsettings line 21")
	project, err := uses.AggregateProjectData(client.Eclient, projID, test1.(string))
	if err != nil {
		fmt.Println(err)
		fmt.Println("error: middleware/project/projectsettings line 23")
	}

	for _, member := range project.ProjectData.Members {
		if member.MemberID == session.Values["DocID"].(string) && member.Role <= 0 {
			cs := client.ClientSide{Project: project}
			client.RenderTemplate(w, "template2-nil", cs)
			client.RenderTemplate(w, "project_settings_F", cs)
			return
		}
	}

	http.Redirect(w, r, "/404/", http.StatusFound)
}
