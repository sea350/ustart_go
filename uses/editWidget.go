package uses

import (
	postWidget "github.com/sea350/ustart_go/post/widget"
	elastic "github.com/olivere/elastic"
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
