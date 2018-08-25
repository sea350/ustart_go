package post

import (
	"context"
	"errors"
	"strings"

	getFollow "github.com/sea350/ustart_go/get/follow"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NewEventFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func NewEventFollow(eclient *elastic.Client, eventID string, field string, newKey string, isBell bool) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.FollowIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	follID, foll, err := getFollow.ByID(eclient, eventID)
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
		if len(foll.EventFollowers) == 0 {
			var newMap = make(map[string]bool)
			newMap[newKey] = isBell
			followMap = newMap
			if isBell {
				var newBell = make(map[string]bool)
				newBell[newKey] = isBell
				bellMap = newBell
			}
		} else {
			foll.EventFollowers[newKey] = isBell
			followMap = foll.UserFollowers

			//modify user bell map if bell follower
			if isBell {
				foll.EventBell[newKey] = isBell
				bellMap = foll.UserBell
			}
		}

	case "following":
		FollowingLock.Lock()
		defer FollowingLock.Unlock()
		if len(foll.EventFollowing) == 0 {
			var newMap = make(map[string]bool)
			newMap[newKey] = isBell
			followMap = newMap
		} else {
			foll.EventFollowing[newKey] = isBell
			followMap = foll.EventFollowing
		}
	default:
		return errors.New("Invalid field")
	}
	var theField string
	if strings.ToLower(field) == "followers" {
		theField = "EventFollowers"
	} else if strings.ToLower(field) == "following" {
		theField = "EventFollowing"
	}
	newFollow := eclient.Update().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		Id(follID).
		Doc(map[string]interface{}{theField: followMap}) //field = Followers or Following, newContent =

	//only executes when there is a new bell follower
	if isBell && strings.ToLower(field) == "followers" {
		newFollow.Doc(map[string]interface{}{"EventBell": bellMap})
	}
	_, err = newFollow.Do(ctx)
	return err
}
