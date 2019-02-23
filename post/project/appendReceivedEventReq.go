package post

import (
	get "github.com/sea350/ustart_go/get/project"
	elastic "github.com/olivere/elastic"
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
