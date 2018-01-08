package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteBlockee ... unblocks a user by deleting from the blocked array
func DeleteBlockee(eclient *elastic.Client, usrID string, blockID string) error {
	ctx := context.Background()

	blockLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := 0
	for i := range usr.BlockedUsers {
		if usr.BlockedUsers[i] == blockID {
			index = i
			break
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	//temp solution end

	usr.BlockedUsers = append(usr.BlockedUsers[:index], usr.BlockedUsers[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"BlockedUsers": usr.BlockedUsers}).
		Do(ctx)

	defer blockLock.Unlock()
	return err

}
