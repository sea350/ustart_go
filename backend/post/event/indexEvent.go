package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/backend/globals"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IndexEvent ... ADDS NEW EVENT TO ES RECORDS
//Needs a type Event struct
//returns the new event's id and an error
func IndexEvent(eclient *elastic.Client, newEvent types.Events) (string, error) {
	//Check if index exists
	ctx := context.Background()
	var ID string
	exists, err := eclient.IndexExists(globals.EventIndex).Do(ctx)
	if err != nil {
		return ID, err
	}
	//If index doesn't exist, create and return it
	if !exists {
		createIndex, err := eclient.CreateIndex(globals.EventIndex).BodyString(globals.MappingEvent).Do(ctx)
		if err != nil {
			_, _ = eclient.IndexExists(globals.EventIndex).Do(ctx)
			panic(err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
		}

		//Return an error saying it doesn't exist
		return ID, errors.New("Index does not exist")
	}

	//Index the document
	newEvnt, Err := eclient.Index().
		Index(globals.EventIndex).
		Type(globals.EventType).
		BodyJson(newEvent).
		Do(ctx)

	if Err != nil {
		return ID, Err
	}

	return newEvnt.Id, nil
}
