package get

import (
	"errors"

	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjAndMember ...
func ProjAndMember(eclient *elastic.Client, projID string, memberID string) (types.Project, types.Member, error) {

	project, err := ProjectByID(eclient, projID)
	if err != nil {
		return types.Project{}, types.Member{}, err
	}

	if len(project.Members) < 1 {
		return types.Project{}, types.Member{}, errors.New("Project has zero members")
	}

	var retMember types.Member
	for mem := range project.Members {
		if project.Members[mem].MemberID == memberID {
			retMember = project.Members[mem]

		}
	}

	return project, retMember, err

}
