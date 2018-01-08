package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendBlocker ... appends to the blocked users array
func AppendBlocker(eclient *elastic.Client, usrID string, blockID string, whichOne bool) error {
	ctx := context.Background()

	BlockLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"BlockedBy": usr.BlockedBy}).
		Do(ctx)

	defer BlockLock.Unlock()
	return err

}
