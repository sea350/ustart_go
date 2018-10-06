package uses

import (
	"fmt"
	"log"

	getEvnt "github.com/sea350/ustart_go/get/event"
	getProj "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"
	postEvnt "github.com/sea350/ustart_go/post/event"
	postProj "github.com/sea350/ustart_go/post/project"
	postUser "github.com/sea350/ustart_go/post/user"
	post "github.com/sea350/ustart_go/post/widget"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AddWidget ...
//Adds a new widget to the UserWidgets array
func AddWidget(eclient *elastic.Client, docID string, newWidget types.Widget, isProject bool, isEvent bool) error {
	if isProject {
		proj, err := getProj.ProjectByID(eclient, docID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return err
		}

		newWidget.Position = len(proj.Widgets)
		widgetID, err := post.IndexWidget(eclient, newWidget)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return err
		}

		updatedWidgets := append(proj.Widgets, widgetID)
		updateErr := postProj.UpdateProject(eclient, docID, "Widgets", updatedWidgets)
		return updateErr
	}

	if isEvent {
		evnt, err := getEvnt.EventByID(eclient, docID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return err
		}

		newWidget.Position = len(evnt.Widgets)
		fmt.Println("newWidget.Position", newWidget.Position)
		widgetID, err := post.IndexWidget(eclient, newWidget)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return err
		}

		updatedWidgets := append(evnt.Widgets, widgetID)
		updateErr := postEvnt.UpdateEvent(eclient, docID, "Widgets", updatedWidgets)
		return updateErr
	}

	usr, err := getUser.UserByID(eclient, docID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}
	newWidget.Position = len(usr.UserWidgets)
	widgetID, err := post.IndexWidget(eclient, newWidget)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	updatedWidgets := append(usr.UserWidgets, widgetID)
	updateErr := postUser.UpdateUser(eclient, docID, "UserWidgets", updatedWidgets)
	return updateErr
}
