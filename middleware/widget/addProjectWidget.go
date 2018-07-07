package widget

import (
	"fmt"
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
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	project, member, err := getProj.ProjAndMember(client.Eclient, r.FormValue("projectWidget"), test1.(string))
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an error: middleware/profile/addProjectWidget.go 26")
	}

	// newWidget, err := ProcessWidgetForm(r)
	// if err != nil {
	// 	fmt.Println("this is an error: middleware/profile/addProjectWidget.go 31")
	// 	fmt.Println(err)
	// 	http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
	// 	return
	// }

	// newWidget.UserID = r.FormValue("projectWidget")

	if uses.HasPrivilege("widget", project.PrivilegeProfiles, member) {

		// if r.FormValue("editID") == `0` {
		// 	fmt.Println("this is debug text middeware/widget/addprojectidget.go")
		// 	fmt.Println(r.FormValue("projectWidget"))
		// err := uses.AddWidget(client.Eclient, r.FormValue("projectWidget"), newWidget, true)
		newWidget, err := ProcessWidgetForm(r)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an error: middleware/profile/addProjectWidget.go 45")
			http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
			return
		}

		newWidget.UserID = r.FormValue("projectWidget")

		if r.FormValue("editID") == `0` {
			// fmt.Println("this is debug text middeware/widget/addprojectidget.go")
			// fmt.Println(r.FormValue("projectWidget"))
			// fmt.Println(newWidget.Data)
			err := uses.AddWidget(client.Eclient, r.FormValue("projectWidget"), newWidget, true, false)
			if err != nil {
				fmt.Println(err)
				fmt.Println("this is an error: middleware/profile/addProjectWidget.go 45")
			}
		} else {
			err := post.ReindexWidget(client.Eclient, r.FormValue("editID"), newWidget)
			if err != nil {
				fmt.Println(err)
				fmt.Println("this is an error: middleware/profile/addWidget.go 51")
			}
		}

		http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
	} else {
		fmt.Println("You do not have the privilege to add a widget to this project. Check your privilege. ")
	}
	return
}
