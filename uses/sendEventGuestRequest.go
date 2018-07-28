package uses

import (
	eventPost "github.com/sea350/ustart_go/post/event"
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//SendEventGuestRequest ...
func SendEventGuestRequest(eclient *elastic.Client, eventID, userID string, classification int) error {
	var err error

	if classification == 1 {
		err := userPost.AppendSentEventReq(eclient, userID, eventID)
		if err != nil {
			return err
		}
	}
	if classification == 2 {
		err := projPost.AppendSentEventReqProject(eclient, userID, eventID)
		if err != nil {
			return err
		}
	}

	//	proj, err := projGet.ProjectByID(eclient, projectInfo.ProjectID)

	if err != nil {
		return err
	}

	err = eventPost.AppendGuestReqReceived(eclient, eventID, userID, classification)
	return err
}
