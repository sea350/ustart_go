package uses

import (
	"time"

	projPost "github.com/sea350/ustart_go/post/project"

	userPost "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AcceptRequest ...
func AcceptRequest(eclient *elastic.Client, projectInfo types.ProjectInfo, userID string) error {
	err := userPost.AppendProject(eclient, userID, projectInfo)
	if err != nil {
		return err
	}
	//	proj, err := projGet.ProjectByID(eclient, projectInfo.ProjectID)

	if err != nil {
		return err
	}

	var newMember types.Member
	newMember.MemberID = userID
	newMember.Role = 2
	newMember.Title = "Member"
	newMember.Visible = true
	newMember.JoinDate = time.Now()

	err = projPost.AppendMember(eclient, projectInfo.ProjectID, newMember)
	return err
}
