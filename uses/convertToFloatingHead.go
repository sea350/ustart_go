package uses

import (
	getProject "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ConvertUserToFloatingHead ... pulls latest version of user and converts relevent data into floating head
func ConvertUserToFloatingHead(eclient *elastic.Client, userDocID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	usr, err := getUser.UserByID(eclient, userDocID)
	if err != nil {
		return head, err
	}

	head.FirstName = usr.FirstName
	head.LastName = usr.LastName
	head.Image = usr.Avatar
	head.Username = usr.Username
	head.DocID = userDocID

	return head, err
}

//ConvertProjectToFloatingHead ... pulls latest version of user and converts relevent data into floating head
func ConvertProjectToFloatingHead(eclient *elastic.Client, projectID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	proj, err := getProject.ProjectByID(eclient, projectID)
	if err != nil {
		panic(err)
	}

	head.FirstName = proj.Name
	head.Bio = proj.Description
	head.Image = proj.Avatar
	head.Username = proj.URLName
	head.DocID = projectID
	head.Notifications = len(proj.MemberReqReceived)

	return head, err
}
