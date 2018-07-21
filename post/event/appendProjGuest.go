package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendProjectGuest ... Appends a new guest to an existing event
//Requires event docID and a types Guest
//Returns an error
func AppendProjectGuest(eclient *elastic.Client, eventID string, guest types.EventProjectGuests) error {
	ctx := context.Background()

	GenericEventUpdateLock.Lock()
	defer GenericEventUpdateLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.ProjectGuests = append(evnt.ProjectGuests, guest)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"ProjectGuests": evnt.ProjectGuests}).
		Do(ctx)

	return err
}
