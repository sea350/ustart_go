package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//AppendProject ... appends new project to user
//takes in eclient, user ID, the project ID, and a bool
func AppendProject(eclient *elastic.Client, usrID string, proj types.ProjectInfo) error {
	ctx := context.Background()

	ProjectLock.Lock()
	defer ProjectLock.Unlock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	isDuplicate := false
	for _, currProj := range usr.Projects {
		if currProj.ProjectID == proj.ProjectID {
			isDuplicate = true
			break
		}
	}

	if !isDuplicate {
		usr.Projects = append(usr.Projects, proj)
	}

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Projects": usr.Projects}).
		Do(ctx)

	return err

}
