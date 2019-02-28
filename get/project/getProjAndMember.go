package get

import (
	"errors"
	"log"

	elastic "github.com/olivere/elastic"
	types "github.com/sea350/ustart_go/types"
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

	for mem := range project.Members {
		log.Println(project.Members[mem].MemberID, "vs.", memberID)

		if project.Members[mem].MemberID == memberID {

			return project, project.Members[mem], err

		}
	}

	return project, types.Member{}, errors.New("Member does not exist")

}
