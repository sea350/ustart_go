package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteEntryID ...deletes entry ID from user array
func DeleteEntryID(eclient *elastic.Client, usrID string, entryID string, idx int) error {
	ctx := context.Background()
	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	usr.EntryIDs = append(usr.EntryIDs[:idx], usr.EntryIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"EntryIDs": usr.EntryIDs}).
		Do(ctx)

	return err

}
