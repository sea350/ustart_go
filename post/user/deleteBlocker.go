package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//DeleteBlocker ... unblocks a user by deleting from the blocked array
func DeleteBlocker(eclient *elastic.Client, usrID string, blockID string, whichOne bool) error {
	ctx := context.Background()

	BlockLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	//temp solution
	index := 0
	for i := range usr.BlockedBy {
		if usr.BlockedBy[i] == blockID {
			index = i
			break
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	//temp solution end
	usr.BlockedBy = append(usr.BlockedBy[:index], usr.BlockedBy[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"BlockedBy": usr.BlockedBy}).
		Do(ctx)

	defer BlockLock.Unlock()
	return err

}
