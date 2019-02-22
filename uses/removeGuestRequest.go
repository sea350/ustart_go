package uses

import (
	evntPost "github.com/sea350/ustart_go/post/event"
	elastic "github.com/olivere/elastic"
)

//RemoveGuestRequest ...
func RemoveGuestRequest(eclient *elastic.Client, eventID string, userID string, classification int) (int, error) {
	var newNumRequests int
	var err error

	newNumRequests, err = evntPost.DeleteGuestReqReceived(eclient, eventID, userID)
	if err != nil {
		return newNumRequests, err
	}

	// if classification == 1 {
	// 	err := userPost.DeleteSentEventReq(eclient, userID, eventID)
	// 	if err != nil {
	// 		return newNumRequests, err
	// 	}
	// }
	// if classification == 2 {
	// 	err := projPost.DeleteSentEventReqProject(eclient, userID, eventID)
	// 	if err != nil {
	// 		return newNumRequests, err
	// 	}
	// }

	return newNumRequests, nil
}
