package uses

import (
	getProject "github.com/sea350/ustart_go/get/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//CreatorRemovesMember ...
//Only project creator can remove anyone from a group
//Admins might be allowed, not sure yet
func CreatorRemovesMember(eclient *elastic.Client, removerID string, removeeID string, projectID string) error {
	proj, err := getProject.ProjectByID(eclient, projectID)

	for i := range proj.Members {
		if proj.Members[i].MemberID == removerID && proj.Members[i].Role == 0 {
			err = RemoveMember(eclient, projectID, removeeID)
			if err != nil {
				panic(err)
			}
		}
	}

	return err

}
