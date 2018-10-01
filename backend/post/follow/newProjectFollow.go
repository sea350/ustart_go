package post

import (
	"context"
	"errors"
	"strings"

	getFollow "github.com/sea350/ustart_go/backend/get/follow"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NewProjectFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func NewProjectFollow(eclient *elastic.Client, projID string, field string, newKey string, isBell bool, followType string) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.FollowIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	follID, foll, err := getFollow.ByID(eclient, projID)
	if err != nil {
		return err
	}

	//  vFollowLock.Lock()

	var followMap = make(map[string]bool)
	var bellMap = make(map[string]bool)
	switch strings.ToLower(field) {
	case "followers":
		FollowLock.Lock()
		defer FollowLock.Unlock()
		if followType == "user" {

			foll.UserFollowers[newKey] = isBell

			followMap = foll.UserFollowers
			//modify user bell map if bell follower
			if isBell {
				bellMap = foll.ProjectBell
				bellMap[newKey] = isBell

			}

		} else if followType == "project" {
			if len(foll.ProjectFollowers) == 0 {
				var newMap = make(map[string]bool)
				newMap[newKey] = isBell
				followMap = newMap
				if isBell {
					var newBell = make(map[string]bool)
					newBell[newKey] = isBell
					bellMap = newBell
				}
			} else {
				foll.ProjectFollowers[newKey] = isBell
				followMap = foll.ProjectFollowers

				//modify user bell map if bell follower
				if isBell {
					foll.ProjectBell[newKey] = isBell
					bellMap = foll.ProjectBell
				}
			}

		}

	case "following":
		FollowLock.Lock()
		defer FollowLock.Unlock()
		if len(foll.ProjectFollowing) == 0 {
			var newMap = make(map[string]bool)
			newMap[newKey] = isBell
			followMap = newMap
		} else {
			foll.ProjectFollowing[newKey] = isBell
			followMap = foll.ProjectFollowing
		}
	default:
		return errors.New("Invalid field")
	}
	newFollow := eclient.Update().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		Id(follID).
		Doc(map[string]interface{}{"UserFollowers": followMap}) //field = Followers or Following, newContent =

	//only executes when there is a new bell follower
	if isBell && strings.ToLower(field) == "followers" {
		newFollow.Doc(map[string]interface{}{"UserBell": bellMap})
	}
	_, err = newFollow.Do(ctx)
	return err
}
