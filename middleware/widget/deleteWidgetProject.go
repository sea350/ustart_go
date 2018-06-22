package widget

import (
	"fmt"
	"net/http"

	getProj "github.com/sea350/ustart_go/get/project"
	get "github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//DeleteWidgetProject ... Deletes a widget and redirects to projects page
func DeleteWidgetProject(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	projectURL := r.FormValue("deleteProjectURL")

	projID, err := getProj.ProjectIDByURL(client.Eclient, projectURL)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error middleware/widget/deleteWidgetProject.go line 29")
	}

	proj, member, err := getProj.ProjAndMember(client.Eclient, projID, test1.(string))

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error middleware/widget/deleteWidgetProject.go line 36")
	}

	if uses.HasPrivilege("widget", proj.PrivilegeProfiles, member) {
		widg, err := get.WidgetByID(client.Eclient, r.FormValue("deleteID"))
		if widg.Classification == 15 {
			http.Redirect(w, r, "/Projects/"+projectURL, http.StatusFound)
			return
		}
		err = uses.RemoveWidget(client.Eclient, r.FormValue("deleteID"), true)
		if err != nil {
			fmt.Println("This is an err, deleteWidgetProject line29")
			fmt.Println(err)
		}

		http.Redirect(w, r, "/Projects/"+projectURL, http.StatusFound)
	} else {
		fmt.Println("You do not have widget privileges for this project.")
	}

	return
}
