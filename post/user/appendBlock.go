package post

import (
	"context"
	"errors"

	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendBlock ... appends to the blocked users array
func AppendBlock(eclient *elastic.Client, usrID string, blockID string, whichOne bool) error {
	ctx := context.Background()

	blockLock.Lock()
	defer blockLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.BlockedUsers = append(usr.BlockedUsers, blockID)

		_, err = eclient.Update().
			Index(globals.UserIndex).
			Type(globals.UserType).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedUsers": usr.BlockedUsers}).
			Do(ctx)

		return err
	} else {
		usr.BlockedUsers = append(usr.BlockedUsers, blockID)

		_, err = eclient.Update().
			Index(globals.UserIndex).
			Type(globals.UserType).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedBy": usr.BlockedBy}).
			Do(ctx)

		return err
	}

}
