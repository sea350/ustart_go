package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendLikedEntryID ... appends to either sent or received project request arrays within user
//takes in eclient, user ID, the project ID, and a bool
//true = append to following, false = append to followers
func AppendLikedEntryID(eclient *elastic.Client, usrID string, entryID string) error {
	ctx := context.Background()

	LikeLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.LikedEntryIDs = append(usr.LikedEntryIDs, entryID)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"LikedEntryIDs": usr.LikedEntryIDs}).
		Do(ctx)

	defer LikeLock.Unlock()
	return err

}
