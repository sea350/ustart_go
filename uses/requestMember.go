package uses

import (
	userPost "github.com/sea350/ustart_go/post/user"

	projPost "github.com/sea350/ustart_go/post/project"
	elastic "github.com/olivere/elastic"
)

//RequestMember ...
func RequestMember(eclient *elastic.Client, projectID string, userID string) error {
	err := userPost.AppendSentProjReq(eclient, userID, projectID)
	if err != nil {
		return err
	}
	err = projPost.AppendMemberReqSent(eclient, projectID, userID)
	return err
}
