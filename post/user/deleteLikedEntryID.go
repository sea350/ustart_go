package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteLikedEntryID ... whichOne: true = following
//whichOne: false = followers
//followID does nothing
func DeleteLikedEntryID(eclient *elastic.Client, usrID string, likerID string) error {
	ctx := context.Background()

	LikeLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := -1
	for i := range usr.LikedEntryIDs {
		if usr.LikedEntryIDs[i] == likerID {
			index = i
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	usr.LikedEntryIDs = append(usr.LikedEntryIDs[:index], usr.LikedEntryIDs[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"LikedEntryIDs": usr.LikedEntryIDs}).
		Do(ctx)

	defer LikeLock.Unlock()
	return err

}
