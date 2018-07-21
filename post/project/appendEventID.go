package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	//	eventPost "github.com/sea350/ustart_go/post/event"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendEventID ... appends a created event ID to project
func AppendEventID(eclient *elastic.Client, projectID string, eventID string) error {
	ctx := context.Background()

	// GenericEventUpdateLock.Lock()
	// defer GenericEventUpdateLock.Unlock()

	proj, err := get.ProjectByID(eclient, projectID)

	if err != nil {
		return errors.New("Project does not exist")
	}

	proj.EventIDs = append(proj.EventIDs, eventID)

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Doc(map[string]interface{}{"EventIDs": proj.EventIDs}).
		Do(ctx)

	return err

}
