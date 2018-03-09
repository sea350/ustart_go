package uses

import (
	"fmt"

	getProj "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"
	postProj "github.com/sea350/ustart_go/post/project"
	postUser "github.com/sea350/ustart_go/post/user"
	post "github.com/sea350/ustart_go/post/widget"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AddWidget ...
//Adds a new widget to the UserWidgets array
func AddWidget(eclient *elastic.Client, docID string, newWidget types.Widget, isProject bool) error {

	if !isProject {
		usr, err := getUser.UserByID(eclient, docID)

		if err != nil {
			panic(err)
		}
		newWidget.Position = len(usr.UserWidgets)
		widgetID, err := post.IndexWidget(eclient, newWidget)
		if err != nil {
			panic(err)
		}

		updatedWidgets := append(usr.UserWidgets, widgetID)
		updateErr := postUser.UpdateUser(eclient, docID, "UserWidgets", updatedWidgets)
		return updateErr
	}
	proj, err := getProj.ProjectByID(eclient, docID)

	if err != nil {
		fmt.Println("this is an err uses/addwidget 37")
		fmt.Println(err)
	}
	newWidget.Position = len(proj.Widgets)
	widgetID, err := post.IndexWidget(eclient, newWidget)
	if err != nil {
		panic(err)
	}
	updatedWidgets := append(proj.Widgets, widgetID)
	updateErr := postProj.UpdateProject(eclient, docID, "Widgets", updatedWidgets)
	return updateErr

}
