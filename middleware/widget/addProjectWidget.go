package widget

import (
	"net/http"

	getProj "github.com/sea350/ustart_go/get/project"

	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/widget"
	"github.com/sea350/ustart_go/uses"
)

//AddProjectWidget ... After widget form submission adds a widget to database
func AddProjectWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()

	project, member, err := getProj.ProjAndMember(client.Eclient, r.FormValue("projectWidget"), test1.(string))
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		if project.URLName != `` {
			http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
		}
		return
	}

	if uses.HasPrivilege("widget", project.PrivilegeProfiles, member) {
		newWidget, err := ProcessWidgetForm(r)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
			return
		}

		newWidget.UserID = r.FormValue("projectWidget")

		if r.FormValue("editID") == `0` {
			err := uses.AddWidget(client.Eclient, r.FormValue("projectWidget"), newWidget, true, false)
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}
		} else {
			err := post.ReindexWidget(client.Eclient, r.FormValue("editID"), newWidget)
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}
		}

		http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
	} else {

		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "You do not have the privilege to add a widget to this project. Check your privilege. ")
	}
	return
}
