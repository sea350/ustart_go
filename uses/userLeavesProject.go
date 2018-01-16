package uses

import (
	getProject "github.com/sea350/ustart_go/get/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserLeavesProject ...
//When a user leaves a project, checks for creator status
//if creator, a new creator/leader must be assigned
func UserLeavesProject(eclient *elastic.Client, leaverID string, projectID string, newLeaderID string) error {
	proj, err := getProject.ProjectByID(eclient, projectID)

	for i := range proj.Members {
		if proj.Members[i].MemberID == leaverID && proj.Members[i].Role == 0 {
			err = NewProjectLeader(eclient, newLeaderID)
			if err != nil {
				panic(err)
			}
		}
	}

	err = RemoveMember(eclient, projectID, leaverID)

	if err != nil {
		panic(err)
	}

	return err

}
