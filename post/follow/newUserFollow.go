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

//NewUserFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func NewUserFollow(eclient *elastic.Client, userID string, field string, newKey string, isBell bool) error {

	ctx := context.Background()
	fmt.Println("USERID NUF:", userID)
	exists, err := eclient.IndexExists(globals.FollowIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	follID, foll, err := getFollow.ByID(eclient, userID)
	fmt.Println("FOLLID,", follID)
	if err != nil {
		return err
	}

	//  vFollowerLock.Lock()

	var followMap = make(map[string]bool)
	var bellMap = make(map[string]bool)
	switch strings.ToLower(field) {
	case "followers":
		fmt.Println("CASE FOLLOWERS, LINE 41")
		FollowerLock.Lock()
		defer FollowerLock.Unlock()
		if len(foll.UserFollowers) == 0 {
			var newMap = make(map[string]bool)
			newMap[newKey] = isBell
			followMap = newMap
			if isBell {
				var newBell = make(map[string]bool)
				newBell[newKey] = isBell
				bellMap = newBell
			}
		} else {
			foll.UserFollowers[newKey] = isBell
			followMap = foll.UserFollowers

			//modify user bell map if bell follower
			if isBell {
				foll.UserBell[newKey] = isBell
				bellMap = foll.UserBell
			}
		}

	case "following":
		FollowingLock.Lock()
		defer FollowingLock.Unlock()
		fmt.Println("CASE FOLLOWING, LINE 67")
		if len(foll.UserFollowing) == 0 {
			fmt.Println("CASE FOLLOWING, LINE 69")
			var newMap = make(map[string]bool)
			newMap[newKey] = isBell
			followMap = newMap
		} else {
			foll.UserFollowing[newKey] = isBell
			followMap = foll.UserFollowing
		}
	default:
		return errors.New("Invalid field")
	}
	fmt.Println("CURRENT FOLLOW MAP:", followMap)

	var theField string
	if strings.ToLower(field) == "followers" {
		theField = "UserFollowers"
	} else if strings.ToLower(field) == "followers" {
		theField = "UserFollowing"
	}
	newFollow := eclient.Update().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		Id(follID).
		Doc(map[string]interface{}{theField: followMap}) //field = Followers or Following, newContent =

	//only executes when there is a new bell follower
	if isBell && strings.ToLower(field) == "followers" {
		newFollow.Doc(map[string]interface{}{"UserBell": bellMap})
	}
	_, err = newFollow.Do(ctx)
	if err != nil {
		fmt.Println("LINE 92 ERROR,", err)
	}
	fmt.Println("LINE 95, S U C C ")
	return err
}
