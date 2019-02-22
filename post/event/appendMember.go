package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//AppendMember ... APPENDS A NEW MEMBER TO AN EXISTING PROJECT DOC
//Requires the project docID and a type Member
//Returns an error
func AppendMember(eclient *elastic.Client, eventID string, member types.EventMembers) error {
	ctx := context.Background()

	GenericEventUpdateLock.Lock()
	defer GenericEventUpdateLock.Unlock()

	EventMemberLock.Lock()
	defer EventMemberLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.Members = append(evnt.Members, member)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"Members": evnt.Members}).
		Do(ctx)

	return err

}
