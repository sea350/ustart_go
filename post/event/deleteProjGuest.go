package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	getProj "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	postProj "github.com/sea350/ustart_go/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteProjectGuest ... REMOVES A SPECIFIC GUEST FROM AN ARRAY
//Requires event docID and a type guest
//Returns an error
func DeleteProjectGuest(eclient *elastic.Client, eventID string, projectID string) error {

	ctx := context.Background()

	EventGuestLock.Lock()
	defer EventGuestLock.Unlock()

	usr, err := getProj.ProjectByID(eclient, projectID)
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
		err = postProj.UpdateProject(eclient, projectID, "Events", append(usr.Events[:usrIdx], usr.Events[usrIdx+1:]...))
	}
	if err != nil {
		return err
	}

	if usrIdx == len(usr.Events) {
		err = postProj.UpdateProject(eclient, projectID, "Events", nil)
	}
	if err != nil {
		return err
	}
	index := -1
	for i := range evnt.ProjectGuests {
		if evnt.ProjectGuests[i].ProjectID == projectID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("Guest not found")
	}

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"ProjectGuests": evnt.ProjectGuests}).
		Do(ctx)

	return err

}
