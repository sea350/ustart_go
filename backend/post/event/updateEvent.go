package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/event"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateEvent ...
func UpdateEvent(eclient *elastic.Client, eventID string, field string, newContent interface{}) error {
	//UPDATES A SINGLE FEILD IN AN EXISTING ES DOC
	//Requires the docID, feild to be modified, and the new content
	//Returns an error

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.EventIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	GenericEventUpdateLock.Lock()
	_, err = get.EventByID(eclient, eventID)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	defer GenericEventUpdateLock.Unlock()
	return err
}
