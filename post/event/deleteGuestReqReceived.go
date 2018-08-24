package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteGuestReqReceived ... DELETES A USER ID IN A EVENT'S GuestReqReceived ARRAY
//Requires event's docID and the user's docID
//Returns an error
func DeleteGuestReqReceived(eclient *elastic.Client, eventID string, userID string) (int, error) {
	var numRequests int
	ctx := context.Background()

	EventGuestRequestLock.Lock()
	defer EventGuestRequestLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return numRequests, errors.New("Event does not exist")
	}

	delete(evnt.GuestReqReceived, userID)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"GuestReqReceived": evnt.GuestReqReceived}).
		Do(ctx)

	return len(evnt.GuestReqReceived), err

}
