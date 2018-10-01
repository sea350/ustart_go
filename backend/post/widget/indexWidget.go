package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/backend/globals"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IndexWidget ...
// adds a new widget document to the ES cluster
// returns err, nil if successful.
func IndexWidget(eclient *elastic.Client, newWidget types.Widget) (string, error) {
	// Check if the index exists
	ctx := context.Background()
	var id string
	exists, err := eclient.IndexExists(globals.WidgetIndex).Do(ctx)
	if err != nil {
		return id, err
	}
	// If the index doesn't exist, create it and return error.
	if !exists {
		createIndex, Err := eclient.CreateIndex(globals.WidgetIndex).BodyString(globals.MappingWidget).Do(ctx)
		if Err != nil {
			_, _ = eclient.IndexExists(globals.WidgetIndex).Do(ctx)
			panic(Err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
		}

		// Return an error saying it doesn't exist
		return id, errors.New("Index does not exist")
	}

	// Index the document.
	newWidg, Err := eclient.Index().
		Index(globals.WidgetIndex).
		Type(globals.WidgetType).
		BodyJson(newWidget).
		Do(ctx)

	if Err != nil {
		return id, Err
	}

	return newWidg.Id, nil
}
