package uses

import (
	eventPost "github.com/sea350/ustart_go/post/event"

	userPost "github.com/sea350/ustart_go/post/user"
	elastic "github.com/olivere/elastic"
)

//SendEventMemberRequest ...
func SendEventMemberRequest(eclient *elastic.Client, eventID, userID string) error {
	err := userPost.AppendSentEventReq(eclient, userID, eventID)
	if err != nil {
		return err
	}
	//	proj, err := projGet.ProjectByID(eclient, projectInfo.ProjectID)

	if err != nil {
		return err
	}

	err = eventPost.AppendMemberReqReceived(eclient, eventID, userID)
	return err
}
