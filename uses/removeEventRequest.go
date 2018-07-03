package uses

import (
	evntPost "github.com/sea350/ustart_go/post/event"
	userPost "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveGuestRequest ...
func RemoveGuestRequest(eclient *elastic.Client, eventID string, userID string) (int, error) {
	var newNumRequests int
	err := userPost.DeleteSentEventReq(eclient, userID, eventID)
	if err != nil {
		return newNumRequests, err
	}
	newNumRequests, err = evntPost.DeleteMemberReqReceived(eclient, eventID, userID)
	return newNumRequests, err
}
