package uses

import (
	userPost "github.com/sea350/ustart_go/post/user"

	projPost "github.com/sea350/ustart_go/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RequestMember ...
func RequestMember(eclient *elastic.Client, projectID string, userID string) error {
	err := userPost.AppendProjReq(eclient, userID, projectID, false)
	if err != nil {
		return err
	}
	err = projPost.AppendMemberReqSent(eclient, projectID, userID)
	return err
}
