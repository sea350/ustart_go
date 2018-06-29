package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendMemberReqSent ... Appends a User ID To Events MemberReqSent Array
func AppendMemberReqSent(eclient *elastic.Client, eventID string, userID string) error {

	ctx := context.Background()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.MemberReqSent = append(evnt.MemberReqSent, userID)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"MemberReqSent": evnt.MemberReqSent}).
		Do(ctx)

	return err
}
