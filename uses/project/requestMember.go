package uses

import elastic "gopkg.in/olivere/elastic.v5"

//RequestMember ...
func RequestMember(eclient *elastic.Client, projectID string, userID string) error {
	err := post.AppendProjReq(eclient, userID, projectID, false)
	if err != nil {
		return err
	}
	err = post.AppendMemberReqSent(eclient, projectID, userID)
	return err
}
