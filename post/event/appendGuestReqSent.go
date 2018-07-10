package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendGuestReqSent ... Appends a User ID To Events GuestReqSent Array
func AppendGuestReqSent(eclient *elastic.Client, eventID string, userID string) error {

	ctx := context.Background()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.GuestReqSent = append(evnt.GuestReqSent, userID)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"GuestReqSent": evnt.GuestReqSent}).
		Do(ctx)

	return err
}