package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ReindexEvent ...
//  Add a new Event to ES.
//  Returns an error, nil if successful
func ReindexEvent(eclient *elastic.Client, eventID string, newEvent types.Events) error {
	ctx := context.Background()

	GenericEventUpdateLock.Lock()
	defer GenericEventUpdateLock.Unlock()

	exists, err := eclient.IndexExists(globals.EventIndex).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		BodyJson(newEvent).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil

}
