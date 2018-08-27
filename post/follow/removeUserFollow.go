package post

import (
	"context"
	"errors"
	"fmt"
	"strings"

	getFollow "github.com/sea350/ustart_go/get/follow"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveUserFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func RemoveUserFollow(eclient *elastic.Client, userID string, field string, deleteKey string) error {

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

	// var followMap = make(map[string]bool)
	switch strings.ToLower(field) {
	case "followers":
		fmt.Println("REMOVING FOLLOWERS")
		FollowerLock.Lock()
		defer FollowerLock.Unlock()
		if len(foll.UserFollowers) == 0 {
			return errors.New("No followers to remove")
		}
		delete(foll.UserFollowers, deleteKey)

	case "following":
		fmt.Println("REMOVING FOLLOWING")
		FollowingLock.Lock()
		defer FollowingLock.Unlock()
		if len(foll.UserFollowing) == 0 {
			return errors.New("Nothing to remove from following")
		}
		delete(foll.UserFollowing, deleteKey)

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
