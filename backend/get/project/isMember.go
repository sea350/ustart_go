package get

import (
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IsMember ...
func IsMember(eclient *elastic.Client, userID string, project types.Project) bool {

	if len(project.Members) < 1 {
		return false
	}

	for mem, _ := range project.Members {
		if project.Members[mem].MemberID == userID {
			return true

		}
	}

	return false

}
