package uses

import (
	evntPost "github.com/sea350/ustart_go/post/event"
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveGuestRequest ...
func RemoveGuestRequest(eclient *elastic.Client, eventID string, userID string, classification int) (int, error) {
	var newNumRequests int
	var err error
	if classification == 1 {
		err := userPost.DeleteSentEventReq(eclient, userID, eventID)
		if err != nil {
			return newNumRequests, err
		}
	}
	if classification == 2 {
		err := projPost.DeleteSentEventReqProject(eclient, userID, eventID)
		if err != nil {
			return newNumRequests, err
		}
	}

	newNumRequests, err = evntPost.DeleteGuestReqReceived(eclient, eventID, userID)
	return newNumRequests, err
}
