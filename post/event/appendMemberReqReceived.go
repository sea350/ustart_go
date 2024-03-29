package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//AppendMemberReqReceived ... APPENDS A USER ID TO A EVENTS MemberReqReceived ARRAY
//Requires event's docID and the user's docID
//Returns an error
func AppendMemberReqReceived(eclient *elastic.Client, eventID string, userID string) error {

	ctx := context.Background()

	GenericEventUpdateLock.Lock()
	defer GenericEventUpdateLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.MemberReqReceived = append(evnt.MemberReqReceived, userID) //retrieve event

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"MemberReqReceived": evnt.MemberReqReceived}).
		Do(ctx)

	return err

}
