package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteSentEventReqProject ... whichOne: true = sent
//whichOne: false = received
func DeleteSentEventReqProject(eclient *elastic.Client, projectID string, eventID string) error {
	ctx := context.Background()

	GenericProjectUpdateLock.Lock()
	defer GenericProjectUpdateLock.Unlock()

	usr, err := get.ProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	index := -1
	for i := range usr.SentEventReq {
		if usr.SentEventReq[i] == eventID {
			index = i
			break
		}
	}

	if index < 0 {
		return errors.New("index does not exist")
	}
	//end of temp solution

	usr.SentEventReq = append(usr.SentEventReq[:index], usr.SentEventReq[index+1:]...)

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Doc(map[string]interface{}{"SentEventReq": usr.SentEventReq}).
		Do(ctx)

	return err

}
