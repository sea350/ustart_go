package post

import (
	"errors"

	get "github.com/sea350/ustart_go/backend/get/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendSentEventReqProject ... appends to either sent event request arrays within user
//takes in eclient, project ID, the event ID
func AppendSentEventReqProject(eclient *elastic.Client, projectID string, eventID string) error {
	GenericProjectUpdateLock.Lock()

	usr, err := get.ProjectByID(eclient, projectID)

	if err != nil {
		return errors.New("Project does not exist")
	}

	usr.SentEventReq = append(usr.SentEventReq, eventID)

	err = UpdateProject(eclient, projectID, "SentEventReq", usr.SentEventReq)

	defer GenericProjectUpdateLock.Unlock()
	return err

}
