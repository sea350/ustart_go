package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/widget"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//UpdateWidget ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
func UpdateWidget(eclient *elastic.Client, widgetID string, field string, newContent interface{}) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.WidgetIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.WidgetByID(eclient, widgetID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.WidgetIndex).
		Type(globals.WidgetType).
		Id(widgetID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	return err
}

// //UpdateWidget ...
// //  Change a single field of the ES Document
// //  Return an error, nil if successful
// func UpdateWidgetData(eclient *elastic.Client, widgetID string, index int) error {

// ctx := context.Background()

// exists, err := eclient.IndexExists(globals.WidgetIndex).Do(ctx)
// if err != nil {
// 	return err
// }
// if !exists {
// 	return errors.New("Index does not exist")
// }

// _, err = get.WidgetByID(eclient, widgetID)
// if err != nil {
// 	return err
// }

// _, err = eclient.Update().
// 	Index(globals.WidgetIndex).
// 	Type(globals.WidgetType).
// 	Id(widgetID).
// 	Doc(map[string]interface{}{field: newContent}).
// 	Do(ctx)

// return err
// }
