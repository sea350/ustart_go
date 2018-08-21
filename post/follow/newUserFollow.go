package post

import (
	"context"
	"errors"
	"strings"

	getFollow "github.com/sea350/ustart_go/get/follow"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NewUserFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func NewUserFollow(eclient *elastic.Client, userID string, field string, newKey string, isBell bool) error {

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

	//  vFollowerLock.Lock()

	var followMap = make(map[string]bool)
	var bellMap = make(map[string]bool)
	switch strings.ToLower(field) {
	case "followers":
		FollowerLock.Lock()
		defer FollowerLock.Unlock()
		foll.UserFollowers[newKey] = isBell
		followMap = foll.UserFollowers
		//modify user bell map if bell follower
		if isBell {
			foll.UserBell[newKey] = isBell
			bellMap = foll.UserBell
		}

	case "following":
		FollowingLock.Lock()
		defer FollowingLock.Unlock()
		foll.UserFollowing[newKey] = isBell
		followMap = foll.UserFollowing
	default:
		return errors.New("Invalid field")
	}
	newFollow := eclient.Update().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		Id(follID).
		Doc(map[string]interface{}{field: followMap}) //field = Followers or Following, newContent =

	//only executes when there is a new bell follower
	if isBell && strings.ToLower(field) == "followers" {
		newFollow.Doc(map[string]interface{}{"UserBell": bellMap})
	}
	_, err = newFollow.Do(ctx)
	return err
}
