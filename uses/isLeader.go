package uses

import (
	get "github.com/sea350/ustart_go/get/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IsLeader ...
//Returns bool to represent whether member is project leader
//ALso returns index of member
func IsLeader(eclient *elastic.Client, projectID string, memberID string) (bool, int) {
	proj, err := get.ProjectByID(eclient, projectID)
	if err != nil {
		panic(err)
	}
	idx := -1
	for i := range proj.Members {
		//idx = i
		if proj.Members[i].MemberID == memberID && proj.Members[i].Role == 0 {
			idx = i
			return true, idx
		}
	}

	return false, idx
}
