package post

import (
	"context"
	"errors"
	"strings"

	getFollow "github.com/sea350/ustart_go/backend/get/follow"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveEventFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func RemoveEventFollow(eclient *elastic.Client, userID string, field string, deleteKey string) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.FollowIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	follID, foll, err := getFollow.ByID(eclient, userID)
	if err != nil {
		return err
	}

	switch strings.ToLower(field) {
	case "followers":
		FollowLock.Lock()
		defer FollowLock.Unlock()
		if len(foll.UserFollowing) == 0 {
			return errors.New("Nothing to remove from followers")
		}
		delete(foll.EventFollowers, deleteKey)

	case "following":
		FollowLock.Lock()
		defer FollowLock.Unlock()
		if len(foll.UserFollowing) == 0 {
			return errors.New("Nothing to remove from following")
		}
		delete(foll.EventFollowing, deleteKey)

	default:
		return errors.New("Invalid field")
	}

	_, err = eclient.Index().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		Id(follID).
		BodyJson(foll). //field = Followers or Following, newContent =
		Do(ctx)

	return err
}
