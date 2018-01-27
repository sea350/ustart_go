package uses

import (
	get "github.com/sea350/ustart_go/get/widget"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//LoadWidgets ... Loads a list of widgets
//Requires an array of widget ids
//Returns an array/slice of the data for those ids as types.Widget, and an error
func LoadWidgets(eclient *elastic.Client, loadList []string) ([]types.Widget, []error) {

	var widgets []types.Widget
	var errStrings []error
	for _, widgetID := range loadList {
		nextWidget, err := get.WidgetByID(eclient, widgetID)

		if err == nil {
			nextWidget.ID = widgetID
			widgets = append(widgets, nextWidget)

		} else {
			errStrings = append(errStrings, err)
		}

	}

	if len(errStrings) > 0 {
		return widgets, errStrings
	}
	return widgets, nil
}
