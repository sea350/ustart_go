package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Project ...
func Project(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	project, err := uses.AggregateProjectData(client.Eclient, r.FormValue("ProjectURL"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("error: middleware/project/projectsettings line 23")
	}

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		fmt.Println(err)
		fmt.Println("error: middleware/project/projectsettings line 29")
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), Project: project}
	client.RenderTemplate(w, "template2-nil", cs)
	client.RenderTemplate(w, "project_settings_F", cs)
}
