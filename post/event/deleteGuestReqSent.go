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

	//replace with universal.FindIndex when it works
	index := -1
	for i := range evnt.GuestReqSent {
		if evnt.GuestReqSent[i] == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("link not found")
	}

	evnt.GuestReqSent = append(evnt.GuestReqSent[:index], evnt.GuestReqSent[index+1:]...)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"GuestReqSent": evnt.GuestReqSent}).
		Do(ctx)

	return err
}
