package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteMemberReqReceived ... DELETES A USER ID IN A EVENT'S MemberReqReceived ARRAY
//Requires event's docID and the user's docID
//Returns an error
func DeleteMemberReqReceived(eclient *elastic.Client, eventID string, userID string) (int, error) {
	var numRequests int
	ctx := context.Background()

	GenericEventUpdateLock.Lock()
	defer GenericEventUpdateLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return numRequests, errors.New("Event does not exist")
	}

	//replace with universal.FindIndex when it works
	index := -1
	for i := range evnt.MemberReqReceived {
		if evnt.MemberReqReceived[i] == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return numRequests, errors.New("link not found")
	}

	evnt.MemberReqReceived = append(evnt.MemberReqReceived[:index], evnt.MemberReqReceived[index+1:]...)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"MemberReqReceived": evnt.MemberReqReceived}).
		Do(ctx)

	return len(evnt.MemberReqReceived), err

}
