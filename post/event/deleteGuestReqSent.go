package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteGuestReqSent ... DELETES A USER ID IN A EVENT'S GuestReqSent ARRAY
//Requires event's docID and the user's docID
//Returns an error
func DeleteGuestReqSent(eclient *elastic.Client, eventID string, userID string) error {
	ctx := context.Background()
	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	GenericEventUpdateLock.Lock()
	defer GenericEventUpdateLock.Unlock()

	delete(evnt.GuestReqSent, userID)

	_, err = eclient.Index().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		BodyJson(evnt).
		Do(ctx)

	return err
}
