package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/project"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteReceivedEventReqProject ...
func DeleteReceivedEventReqProject(eclient *elastic.Client, projectID string, eventID string) error {
	ctx := context.Background()

	GenericProjectUpdateLock.Lock()

	usr, err := get.ProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	index := 0
	for i := range usr.ReceivedEventReq {
		if usr.ReceivedEventReq[i] == eventID {
			index = i
			break
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	//end of temp solution
	usr.ReceivedEventReq = append(usr.ReceivedEventReq[:index], usr.ReceivedEventReq[index+1:]...)

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Doc(map[string]interface{}{"ReceivedEventReq": usr.ReceivedEventReq}).
		Do(ctx)

	defer GenericProjectUpdateLock.Unlock()
	return err
}
