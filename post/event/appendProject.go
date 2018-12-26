package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendProject ... Appends project to existing event
//Requires the project docID and a type EventProjects
//Returns an error
func AppendProject(eclient *elastic.Client, eventID string, project types.EventProjects) error {
	ctx := context.Background()

	GenericEventUpdateLock.Lock()
	defer GenericEventUpdateLock.Unlock()

	EventMemberLock.Lock()
	defer EventMemberLock.Unlock()

	evnt, err := get.EventByID(eclient, eventID)
	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.Projects = append(evnt.Projects, project)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(eventID).
		Doc(map[string]interface{}{"Projects": evnt.Projects}).
		Do(ctx)

	return err

}
