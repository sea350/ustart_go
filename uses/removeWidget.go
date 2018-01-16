package uses

import (
	getUser "github.com/sea350/ustart_go/get/user"
	getWidget "github.com/sea350/ustart_go/get/widget"
	postUser "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
	"context"
)

//RemoveWidget ...
//Removes widget ID from UserWidgets array/slice and widget struct from ES
func RemoveWidget(eclient *elastic.Client, widgetID string) error{
	ctx := context.Background()

	//get widget to use its data
	widget, err := getWidget.WidgetByID(eclient, widgetID)
	userID := widget.UserID

	//get the user
	usr,err := getUser.UserByID(eclient, userID)

	if err!=nil{
		panic(err)
	}

	//update the user widgets array
	updatedWidgets := append(usr.UserWidgets[:widget.Position], usr.UserWidgets[widget.Position+1:]...))
	updateErr:= postUser.UpdateUser(eclient, userID, "UserWidgets", updatedWidgets)

	if updateErr!=nil{
		panic(updateErr)
	}

	//delete the widget from ES
	_, err := eclient.Delete().
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