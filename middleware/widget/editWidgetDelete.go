package widget

import (
	"html/template"
	
	"net/http"
	
	"strings"

	get "github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//EditWidgetDataDelete ... deletes a link in a class 4 or 5 widget widget
func EditWidgetDataDelete(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	username := test1.(string)

	deletedURL := r.FormValue("deleteURL")
	projectURL := r.FormValue("editProjectURL")
	eventURL := r.FormValue("editEventURL")
	var isProject = false
	if projectURL != `` {
		isProject = true
	}

	var isEvent = false
	if eventURL != `` {
		isEvent = true
	}

	oldWidget, err := get.WidgetByID(client.Eclient, r.FormValue("editID"))
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	if len(oldWidget.Data) == 1 && (oldWidget.Classification != 15 && oldWidget.Classification != 16) {
		err = uses.RemoveWidget(client.Eclient, r.FormValue("editID"), isProject, isEvent)
		if err != nil {
			
	
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		if isProject {
			http.Redirect(w, r, "/Projects/"+projectURL, http.StatusFound)
			return
		}
		if isEvent {
			http.Redirect(w, r, "/Events/"+eventURL, http.StatusFound)
			return
		}
		http.Redirect(w, r, "/profile/"+username, http.StatusFound)
		return
	}

	target := -1
	for index, link := range oldWidget.Data {
		if strings.Contains(string(link), deletedURL) || strings.Contains(deletedURL, string(link)) {
			target = index
			break
		}
	}

	var newArr []template.HTML
	if target == -1 {
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Error: middleware/widget/editWidgetDelete line 61 - deleted object not found")
		newArr = oldWidget.Data
	} else if (target + 1) < len(oldWidget.Data) {
		newArr = append(oldWidget.Data[:target], oldWidget.Data[(target+1):]...)
	} else {
		newArr = oldWidget.Data[:target]
	}

	err = uses.EditWidget(client.Eclient, r.FormValue("editID"), newArr)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	if isProject {
		http.Redirect(w, r, "/Projects/"+projectURL, http.StatusFound)
		return
	}
	if isEvent {
		http.Redirect(w, r, "/Events/"+eventURL, http.StatusFound)
		return
	}
	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
	return
}
