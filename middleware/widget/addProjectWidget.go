package widget

import (
	"log"
	"net/http"

	getProj "github.com/sea350/ustart_go/get/project"

	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/widget"
	"github.com/sea350/ustart_go/uses"
)

//AddProjectWidget ... After widget form submission adds a widget to database
func AddProjectWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	project, member, err := getProj.ProjAndMember(client.Eclient, r.FormValue("projectWidget"), test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	if uses.HasPrivilege("widget", project.PrivilegeProfiles, member) {
		newWidget, err := ProcessWidgetForm(r)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
			return
		}

		newWidget.UserID = r.FormValue("projectWidget")

		if r.FormValue("editID") == `0` {
			err := uses.AddWidget(client.Eclient, r.FormValue("projectWidget"), newWidget, true, false)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		} else {
			err := post.ReindexWidget(client.Eclient, r.FormValue("editID"), newWidget)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}

		http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
	} else {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("You do not have the privilege to add a widget to this project. Check your privilege. ")
	}
	return
}
