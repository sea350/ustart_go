package uses

import (
	getUser "github.com/sea350/ustart_go/get/user"
	getWidget "github.com/sea350/ustart_go/get/widget"
	postUser "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AddWidget ...
//Adds a new widget to the UserWidgets array
func AddWidget(eclient *elastic.Client, userID string, newWidget types.Widget) error {
	usr, err := getUser.UserByID(eclient, userID)

	if err != nil {
		panic(err)
	}

	widgetID, err := getWidget.WidgetIDByLink(eclient, newWidget.Link)
	updatedWidgets := append(usr.UserWidgets, widgetID)
	updateErr := postUser.UpdateUser(eclient, userID, "UserWidgets", updatedWidgets)

	return updateErr

}
