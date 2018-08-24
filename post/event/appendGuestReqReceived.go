package post

import (
	"context"
	"errors"
	"fmt"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//Do the same for guests

//AppendGuestReqReceived ... APPENDS A USER ID TO A EVENTS GuestReqReceived ARRAY
//Requires event's docID and the user's docID
//Returns an error
func AppendGuestReqReceived(eclient *elastic.Client, eventID string, userID string, classification int) error {

	ctx := context.Background()

	EventGuestRequestLock.Lock()
	defer EventGuestRequestLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	println("guestreqreceived 0?", evnt.GuestReqReceived)
	if len(evnt.GuestReqReceived) == 0 {
		newReq := make(map[string]int)
		newReq[userID] = classification
		fmt.Println(newReq[userID])
		evnt.GuestReqReceived = newReq

	} else {
		fmt.Println("guestrequest not 0", evnt.GuestReqReceived[userID])
		evnt.GuestReqReceived[userID] = classification
	}

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"GuestReqReceived": evnt.GuestReqReceived}).
		Do(ctx)

	return err

}
