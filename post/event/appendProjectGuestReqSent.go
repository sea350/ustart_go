package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendProjectGuestReqSent ... Appends a User ID To Events ProjectGuestReqSent Array
func AppendProjectGuestReqSent(eclient *elastic.Client, eventID string, projectID string) error {

	ctx := context.Background()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.ProjectGuestReqSent = append(evnt.ProjectGuestReqSent, projectID)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"ProjectGuestReqSent": evnt.ProjectGuestReqSent}).
		Do(ctx)

	return err
}
