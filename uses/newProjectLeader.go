package uses

import (
	get "github.com/sea350/ustart_go/get/project"
	post "github.com/sea350/ustart_go/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NewProjectLeader ...
//Enables current leader to set a new leader
func NewProjectLeader(eclient *elastic.Client, currentLeaderID string, projectID string, newLeaderID string) error {
	proj, err := get.ProjectByID(eclient, projectID)
	if err != nil {
		panic(err)
	}
	isLeader, idx := IsLeader(eclient, projectID, currentLeaderID)

	for i := range proj.Members {
		if proj.Members[i].MemberID == newLeaderID && isLeader {
			proj.Members[i].Role = 0
			proj.Members[idx].Role = 1 //can be any role value, really

		}
	}

	updateErr := post.UpdateProject(eclient, projectID, "Members", proj.Members)

	return updateErr

}
