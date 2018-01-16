package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

const widgetMapping = `
{
    "mappings":{
        "User":{
            "properties":{
                "UserID":{
                    "type":"keyword"
                },
                "Link":{
                	"type":"keyword"
                },
				"Classification":{
					"type":"keyword"
				}

                
            }
        }
    }
}`

//IndexWidget ...
// adds a new widget document to the ES cluster
// returns err, nil if successful.
func IndexUser(eclient *elastic.Client, newWidget types.Widget) error {
	// Check if the index exists
	ctx := context.Background()
	exists, err := eclient.IndexExists(globals.WidgetIndex).Do(ctx)
	if err != nil {
		return err
	}
	// If the index doesn't exist, create it and return error.
	if !exists {
		createIndex, Err := eclient.CreateIndex(globals.WidgetIndex).BodyString(widgetMapping).Do(ctx)
		if Err != nil {
			_, _ = eclient.IndexExists(globals.WidgetIndex).Do(ctx)
			panic(Err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
		}

		// Return an error saying it doesn't exist
		return errors.New("Index does not exist")
	}

	// Index the document.
	_, Err := eclient.Index().
		Index(globals.WidgetIndex).
		Type(globals.WidgetType).
		BodyJson(newWidget).
		Do(ctx)

	if Err != nil {
		return Err
	}

	return nil
}
