package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//Do the same for guests

//AppendProjectGuestReqReceived ... APPENDS A USER ID TO A EVENTS GuestReqReceived ARRAY
//Requires event's docID and the user's docID
//Returns an error
func AppendProjectGuestReqReceived(eclient *elastic.Client, eventID string, projectID string) error {

	ctx := context.Background()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.ProjectGuestReqReceived = append(evnt.ProjectGuestReqReceived, projectID)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"ProjectGuestReqReceived": evnt.ProjectGuestReqReceived}).
		Do(ctx)

	return err

}
