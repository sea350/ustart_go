package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ReindexWidget ...
//  Add a new user to ES.
//  Returns an error, nil if successful
func ReindexWidget(eclient *elastic.Client, widgetID string, widget types.Widget) error {

	ctx := context.Background()
	exists, err := eclient.IndexExists(globals.WidgetIndex).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.WidgetIndex).
		Type(globals.WidgetType).
		Id(widgetID).
		BodyJson(widget).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
