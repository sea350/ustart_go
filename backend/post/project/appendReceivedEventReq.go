package post

import (
	get "github.com/sea350/ustart_go/backend/get/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendReceivedEventReqProject ... appends to either sent or received event request arrays within user
//takes in eclient, user ID, the event ID
func AppendReceivedEventReqProject(eclient *elastic.Client, projectID string, eventID string) error {

	GenericProjectUpdateLock.Lock()

	usr, err := get.ProjectByID(eclient, projectID)

	usr.ReceivedEventReq = append(usr.ReceivedEventReq, eventID)

	err = UpdateProject(eclient, projectID, "ReceivedEventReq", usr.ReceivedEventReq)

	defer GenericProjectUpdateLock.Unlock()
	return err
}
