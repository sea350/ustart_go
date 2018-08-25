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

	var followMap = make(map[string]bool)
	switch strings.ToLower(field) {
	case "followers":
		fmt.Println("REMOVING FOLLOWERS")
		FollowerLock.Lock()
		defer FollowerLock.Unlock()
		delete(foll.UserFollowers, deleteKey)
		followMap = foll.UserFollowers

	case "following":
		fmt.Println("REMOVING FOLLOWING")
		FollowingLock.Lock()
		defer FollowingLock.Unlock()
		delete(foll.UserFollowing, deleteKey)
		followMap = foll.UserFollowing
	default:
		return errors.New("Invalid field")
	}

	// var theField string
	// if strings.ToLower(field) == "followers" {
	// 	theField = "UserFollowers"
	// } else if strings.ToLower(field) == "following" {
	// 	theField = "UserFollowing"
	// }
	// fmt.Println("THE FIELD:", theField)
	// fmt.Println("FOLLOW MAP:", followMap)

	_, err = eclient.Update().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		Id(follID).
		Doc(map[string]interface{}{field: followMap}). //field = Followers or Following, newContent =
		Do(ctx)

	return err
}
