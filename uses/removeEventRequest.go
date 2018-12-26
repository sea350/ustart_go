package uses

import (
	evntPost "github.com/sea350/ustart_go/post/event"
	elastic "gopkg.in/olivere/elastic.v5"
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
