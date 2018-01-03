package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendEntryID ... appends a created entry ID to user
func AppendEntryID(eclient *elastic.Client, usrID string, entryID string) error {
	ctx := context.Background()

	entryLock.Lock()
	defer entryLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.EntryIDs = append(usr.EntryIDs, entryID)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"EntryIDs": usr.EntryIDs}).
		Do(ctx)

	return err

}
