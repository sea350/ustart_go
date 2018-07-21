package uses

import (
	evntPost "github.com/sea350/ustart_go/post/event"
	projPost "github.com/sea350/ustart_go/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveProjectEventRequest ...
func RemoveProjectEventRequest(eclient *elastic.Client, eventID string, projectID string) (int, error) {
	var newNumRequests int
	err := projPost.DeleteSentEventReqProject(eclient, projectID, eventID)
	if err != nil {
		return newNumRequests, err
	}
	newNumRequests, err = evntPost.DeleteProjectGuestReqReceived(eclient, eventID, projectID)
	return newNumRequests, err
}
