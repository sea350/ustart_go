package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//DeleteMemberReqSent ... DELETES A USER ID IN A EVENT'S MemberReqSent ARRAY
//Requires event's docID and the user's docID
//Returns an error
func DeleteMemberReqSent(eclient *elastic.Client, eventID string, userID string) error {
	ctx := context.Background()

	GenericEventUpdateLock.Lock()
	defer GenericEventUpdateLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	//replace with universal.FindIndex when it works
	index := -1
	for i := range evnt.MemberReqSent {
		if evnt.MemberReqSent[i] == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("link not found")
	}

	evnt.MemberReqSent = append(evnt.MemberReqSent[:index], evnt.MemberReqSent[index+1:]...)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"MemberReqSent": evnt.MemberReqSent}).
		Do(ctx)

	return err
}
