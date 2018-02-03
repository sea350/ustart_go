package uses

import (
	"context"

	getUser "github.com/sea350/ustart_go/get/user"
	getWidget "github.com/sea350/ustart_go/get/widget"
	globals "github.com/sea350/ustart_go/globals"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveWidget ...
//Removes widget ID from UserWidgets array/slice and widget struct from ES
func RemoveWidget(eclient *elastic.Client, widgetID string) error {
	ctx := context.Background()

	//get widget to use its data
	widget, err := getWidget.WidgetByID(eclient, widgetID)
	userID := widget.UserID

	//get the user
	usr, err := getUser.UserByID(eclient, userID)

	if err != nil {
		panic(err)
	}

	var pos int
	var updatedWidgets []string
	for index, ID := range usr.UserWidgets {
		if ID == widgetID {
			pos = index
			break
		}
	}
	//update the user widgets array
	if pos+1 < len(usr.UserWidgets) {
		updatedWidgets = append(usr.UserWidgets[:pos], usr.UserWidgets[pos+1:]...)
	} else {
		updatedWidgets = usr.UserWidgets[:pos]
	}

	updateErr := postUser.UpdateUser(eclient, userID, "UserWidgets", updatedWidgets)

	if updateErr != nil {
		panic(updateErr)
	}

	//delete the widget from ES
	_, err = eclient.Delete().
		Index(globals.WidgetType).
		Type(globals.WidgetType).
		Id(widgetID).
		Do(ctx)

	if err != nil {
		// Handle error
		panic(err)
	}

	return err
}
