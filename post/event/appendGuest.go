package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendGuest ... Appends a new guest to an existing event
//Requires event docID and a types Guest
//Returns an error
func AppendGuest(eclient *elastic.Client, eventID string, guest types.EventGuests) error {
	ctx := context.Background()

	EventGuestLock.Lock()
	defer EventGuestLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.Guests = append(evnt.Guests, guest)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"Guests": evnt.Guests}).
		Do(ctx)

	return err
}
