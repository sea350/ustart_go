package uses

import (
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
	elastic "github.com/olivere/elastic"
)

//RemoveRequest ...
func RemoveRequest(eclient *elastic.Client, projectID string, userID string) (int, error) {
	var newNumRequests int
	err := userPost.DeleteSentProjReq(eclient, userID, projectID)
	if err != nil {
		return newNumRequests, err
	}
	newNumRequests, err = projPost.DeleteMemberReqReceived(eclient, projectID, userID)
	return newNumRequests, err
}
