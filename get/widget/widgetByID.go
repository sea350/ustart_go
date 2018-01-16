package get

import (
	"context"
	"encoding/json"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
	//post "github.com/sea350/ustart_go/post"
)

//WidgetByID ...
func WidgetByID(eclient *elastic.Client, widgetID string) (types.Widget, error) {
	ctx := context.Background()         //intialize context background
	var widget types.Widget             //initialize type widget
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.UserIndex).
						Type(globals.UserType).
						Id(widgetID).
						Do(ctx)

	if err != nil {
		return widget, err
	}

	Err := json.Unmarshal(*searchResult.Source, &widget) //unmarshal type RawMessage into widget struct

	return widget, Err

}
