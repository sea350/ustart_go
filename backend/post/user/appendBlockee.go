package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendBlockee ... puts user on a blocked list
func AppendBlockee(eclient *elastic.Client, usrID string, blockID string) error {
	ctx := context.Background()

	BlockLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.BlockedUsers = append(usr.BlockedUsers, blockID)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"BlockedUsers": usr.BlockedUsers}).
		Do(ctx)

	defer BlockLock.Unlock()
	return err

}
