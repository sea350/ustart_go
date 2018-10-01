package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/event"
	getUser "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	postUser "github.com/sea350/ustart_go/backend/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteMember ... REMOVES A SPECIFIC MEMBER FROM AN ARRAY
//Requires event docID and a type member
//Returns an error
func DeleteMember(eclient *elastic.Client, eventID string, userID string) error {

	ctx := context.Background()

	EventMemberLock.Lock()
	defer EventMemberLock.Unlock()

	usr, err := getUser.UserByID(eclient, userID)
	if err != nil {
		return err
	}
	evnt, evntErr := get.EventByID(eclient, eventID)
	if evntErr != nil {
		return evntErr
	}

	var usrIdx int

	for idx := range usr.Events {
		if usr.Events[idx].EventID == eventID {
			usrIdx = idx
			break
		}
	}

	if usrIdx < len(usr.Events)-1 {
		err = postUser.UpdateUser(eclient, userID, "Events", append(usr.Events[:usrIdx], usr.Events[usrIdx+1:]...))
	}
	if err != nil {
		return err
	}

	if usrIdx == len(usr.Events) {
		err = postUser.UpdateUser(eclient, userID, "Events", nil)
	}
	if err != nil {
		return err
	}

	index := -1
	for i := range evnt.Members {
		if evnt.Members[i].MemberID == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("Member not found")
	}

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"Members": evnt.Members}).
		Do(ctx)

	return err
}
