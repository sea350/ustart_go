package uses

import (
	"context"
	"fmt"

	getEvnt "github.com/sea350/ustart_go/get/event"
	getProj "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"
	getWidget "github.com/sea350/ustart_go/get/widget"
	globals "github.com/sea350/ustart_go/globals"
	postEvnt "github.com/sea350/ustart_go/post/event"
	postProj "github.com/sea350/ustart_go/post/project"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveWidget ...
//Removes widget ID from UserWidgets array/slice and widget struct from ES
func RemoveWidget(eclient *elastic.Client, widgetID string, isProject bool, isEvent bool) error {
	ctx := context.Background()

	//get widget to use its data
	widget, err := getWidget.WidgetByID(eclient, widgetID)
	userID := widget.UserID

	//get the user
	var oldArray []string
	if isProject {
		proj, err := getProj.ProjectByID(eclient, userID)
		if err != nil {
			panic(err)
		}
		oldArray = proj.Widgets
	} else if isEvent {
		event, err := getEvnt.EventByID(eclient, userID)
		if err != nil {
			panic(err)
		}
		oldArray = event.Widgets
	} else {
		usr, err := getUser.UserByID(eclient, userID)
		if err != nil {
			panic(err)
		}
		oldArray = usr.UserWidgets
	}

	var pos int
	var updatedWidgets []string
	for index, ID := range oldArray {
		if ID == widgetID {
			pos = index
			break
		}
	}
	//update the user widgets array
	if pos+1 < len(oldArray) {
		updatedWidgets = append(oldArray[:pos], oldArray[pos+1:]...)
	} else {
		updatedWidgets = oldArray[:pos]
	}

	if isProject {
		updateErr := postProj.UpdateProject(eclient, userID, "Widgets", updatedWidgets)

		if updateErr != nil {
			fmt.Println(updateErr)
			fmt.Println("this is an error, uses/removeWidget line 50")
		}
	} else if isEvent {
		updateErr := postEvnt.UpdateEvent(eclient, userID, "Widgets", updatedWidgets)
		if updateErr != nil {
			fmt.Println(updateErr)
			fmt.Println("this is an error, uses/removeWidget line 50")
		}
	} else {
		updateErr := postUser.UpdateUser(eclient, userID, "UserWidgets", updatedWidgets)

		if updateErr != nil {
			fmt.Println(updateErr)
			fmt.Println("this is an error, uses/removeWidget line 57")
		}
	}

	//delete the widget from ES
	_, err = eclient.Delete().
		Index(globals.WidgetIndex).
		Type(globals.WidgetType).
		Id(widgetID).
		Do(ctx)

	if err != nil {
		// Handle error
		panic(err)
	}

	return err
}
