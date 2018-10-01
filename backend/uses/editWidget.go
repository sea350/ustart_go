package uses

import (
	postWidget "github.com/sea350/ustart_go/backend/post/widget"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EditWidget ...
//Edits an existing widget in the UserWidgets array
func EditWidget(eclient *elastic.Client, widgetID string, newVal interface{}) error {
	updateErr := postWidget.UpdateWidget(eclient, widgetID, "Data", newVal)

	if updateErr != nil {
		panic(updateErr)
	}

	return updateErr

}
