package widget

import (
	"encoding/json"
	
	"net/http"

	getProj "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/uses"
)

//SortProjectWidgets ... gets new array of widget ids from project page and updates project struct in ES
func SortProjectWidgets(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()
	sortedWidgets := r.FormValue("sortedWidgets")
	projectURL := r.FormValue("pageID")
	if projectURL == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Project URL not passed in")
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	defer http.Redirect(w, r, "/Projects/"+projectURL, http.StatusFound)

	arr := []string{}
	err := json.Unmarshal([]byte(sortedWidgets), &arr)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	id, err := getProj.ProjectIDByURL(client.Eclient, projectURL)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	project, member, err := getProj.ProjAndMember(client.Eclient, id, docID.(string))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	if uses.HasPrivilege("widget", project.PrivilegeProfiles, member) {
		err = post.UpdateProject(client.Eclient, id, "Widgets", arr)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
	} else {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"You do not have the privilege to add a widget to this project. Check your privilege.")
	}
}
