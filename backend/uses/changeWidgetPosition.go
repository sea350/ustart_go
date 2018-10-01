package uses

import (
	"errors"

	getUser "github.com/sea350/ustart_go/backend/get/user"
	getWidget "github.com/sea350/ustart_go/backend/get/widget"
	postUser "github.com/sea350/ustart_go/backend/post/user"
	postWidget "github.com/sea350/ustart_go/backend/post/widget"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeWidgetPosition ...
//Changes position of a specific widget, then all relative to it
func ChangeWidgetPosition(eclient *elastic.Client, userID string, oldPos int, newPos int) error {
	usr, err := getUser.UserByID(eclient, userID) //retreive user

	if err != nil {
		dneErr := errors.New("User does not exist") //if user dne
		panic(dneErr)
	}

	widgets := usr.UserWidgets //get the widget slice

	movedWidget := widgets[oldPos] //get the widget being moved

	/*widgetID, getErr := getWidget.WidgetIDByLink(eclient, movedWidget.Link)

	if getErr != nil {
		dneErr := errors.New("Widget does not exist")
		panic(dneErr)
	}*/

	if newPos > len(widgets) || newPos < 0 { //is our new position even feasible?
		boundsErr := errors.New("Out of bounds")
		panic(boundsErr)
	}

	if newPos > oldPos { //if the new position is later in the array, move everything within the bounds of the new and old positions back one spot
		for i := oldPos; i < newPos; i++ {
			currWidget, err := getWidget.WidgetByID(eclient, widgets[i])
			if err != nil {
				panic(err)
			}
			updateErr := postWidget.UpdateWidget(eclient, widgets[i], "Position", currWidget.Position-1)

			if updateErr != nil {
				panic(updateErr)
			}
			widgets[i+1] = widgets[i-1]
		}
		currWidget, err := getWidget.WidgetByID(eclient, movedWidget)
		if err != nil {
			panic(nil)
		}
		currWidget.Position = newPos
		widgets[newPos] = movedWidget
	} else if newPos < oldPos { //if new position is earlier, move everything in the range forward
		for i := oldPos; i < newPos; i-- {
			currWidget, err := getWidget.WidgetByID(eclient, widgets[i])
			if err != nil {
				panic(err)
			}
			updateErr := postWidget.UpdateWidget(eclient, widgets[i], "Position", currWidget.Position-1)
			if updateErr != nil {
				panic(updateErr)
			}
			widgets[i-1] = widgets[i+1]
		}
		currWidget, err := getWidget.WidgetByID(eclient, movedWidget)
		if err != nil {
			panic(nil)
		}
		currWidget.Position = newPos
		widgets[newPos] = movedWidget
	}

	updateErr := postUser.UpdateUser(eclient, userID, "UserWidgets", widgets)

	return updateErr
}
