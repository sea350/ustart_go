package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendProject ... appends new project to user
//takes in eclient, user ID, the project ID, and a bool
func AppendProject(eclient *elastic.Client, usrID string, proj types.ProjectInfo) error {
	ctx := context.Background()

	projectLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Projects = append(usr.Projects, proj)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Projects": usr.Projects}).
		Do(ctx)

	defer projectLock.Unlock()
	return err

}
