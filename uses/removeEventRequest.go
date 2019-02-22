package uses

import (
	evntPost "github.com/sea350/ustart_go/post/event"
	elastic "github.com/olivere/elastic"
)

//RemoveEventRequest ...
func RemoveEventRequest(eclient *elastic.Client, eventID string, userID string) (int, error) {
	var newNumRequests int

	newNumRequests, err := evntPost.DeleteMemberReqReceived(eclient, eventID, userID)
	if err != nil {
		return newNumRequests, err
	}

	//err = userPost.DeleteSentEventReq(eclient, userID, eventID)
	return newNumRequests, nil
}
