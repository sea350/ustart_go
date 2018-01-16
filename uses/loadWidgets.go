package uses

import (
	get "github.com/sea350/ustart_go/get/widget"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//LoadWidgetd ... Loads a list of widgets
//Requires an array of widget ids
//Returns an array/slice of the data for those ids as types.Widget, and an error
func LoadWidgets(eclient *elastic.Client, loadList []string) ([]types.Widget, error) {

	var widgets []types.Widget

	for _, widgetID := range loadList {
		nextWidget, err := get.WidgetByID(eclient, widgetID)

		if err == nil {
			widgets = append(widgets, nextWidget)

		}

	}

	return widgets, nil
}
