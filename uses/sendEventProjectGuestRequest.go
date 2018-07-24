package uses

import (
	eventPost "github.com/sea350/ustart_go/post/event"

	projPost "github.com/sea350/ustart_go/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//SendEventProjectGuestRequest ...
func SendEventProjectGuestRequest(eclient *elastic.Client, eventID, projectID string) error {
	err := projPost.AppendSentEventReqProject(eclient, projectID, eventID)
	if err != nil {
		return err
	}
	//	proj, err := projGet.ProjectByID(eclient, projectInfo.ProjectID)

	if err != nil {
		return err
	}

	err = eventPost.AppendProjectGuestReqReceived(eclient, eventID, projectID)
	return err
}
