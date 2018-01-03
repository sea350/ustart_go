package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteBlock ... unblocks a user by deleting from the blocked array
func DeleteBlock(eclient *elastic.Client, usrID string, blockID string, whichOne bool) error {
	ctx := context.Background()

	blockLock.Lock()
	defer blockLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		//idx, err := universal.FindIndex(usr.BlockedUsers, blockID)
		//temp solution
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

		return err
	} else {
		//idx, err := universal.FindIndex(usr.BlockedBy, blockID)
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

		return err

	}

}
