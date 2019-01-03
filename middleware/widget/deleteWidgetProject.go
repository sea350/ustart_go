package widget

import (
	
	"net/http"

	getProj "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//DeleteWidgetProject ... Deletes a widget and redirects to projects page
func DeleteWidgetProject(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()
	projectURL := r.FormValue("deleteProjectURL")
	widgetID := r.FormValue("deleteID")
	if projectURL == `` || widgetID == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Critical information was not passed in.")
		return
	}

	defer http.Redirect(w, r, "/Projects/"+projectURL, http.StatusFound)

	id, err := getProj.ProjectIDByURL(client.Eclient, projectURL)

	project, member, err := getProj.ProjAndMember(client.Eclient, id, test1.(string))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		return
	}

	if uses.HasPrivilege("widget", project.PrivilegeProfiles, member) {
		_, err := get.WidgetByID(client.Eclient, widgetID)
		err = uses.RemoveWidget(client.Eclient, widgetID, true, false)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}
	} else {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"You do not have the privilege to add a widget to this project. Check your privilege.")
	}

}
