package uses

import (
	projPost "github.com/sea350/ustart_go/post/project"

	userPost "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//SendProjectRequest ...
func SendProjectRequest(eclient *elastic.Client, projID, userID string) error {
	err := userPost.AppendSentProjReq(eclient, userID, projID)
	if err != nil {
		return err
	}
	//	proj, err := projGet.ProjectByID(eclient, projectInfo.ProjectID)

	if err != nil {
		return err
	}

	err = projPost.AppendMemberReqReceived(eclient, projID, userID)
	return err
}
