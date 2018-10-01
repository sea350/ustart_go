package post

import (
	"context"
	"errors"
	"log"
	"strings"

	getFollow "github.com/sea350/ustart_go/backend/get/follow"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NewUserFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func NewUserFollow(eclient *elastic.Client, userID string, field string, newKey string, isBell bool, followType string) error {

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

	//  vFollowLock.Lock()

	var followMap = make(map[string]bool)
	var bellMap = make(map[string]bool)
	switch strings.ToLower(field) {
	case "followers":

		FollowLock.Lock()
		defer FollowLock.Unlock()
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
		FollowLock.Lock()
		defer FollowLock.Unlock()

		if followType == "user" {
			if len(foll.UserFollowing) == 0 {

				var newMap = make(map[string]bool)
				newMap[newKey] = isBell
				followMap = newMap
			} else {
				foll.UserFollowing[newKey] = isBell
				followMap = foll.UserFollowing
			}
		} else if followType == "project" {
			foll.ProjectFollowing[newKey] = isBell
			followMap = foll.ProjectFollowing
		}
	default:
		return errors.New("Invalid field")
	}

	var theField string
	if strings.ToLower(field) == "followers" {
		if followType == "user" {
			theField = "UserFollowers"
		} else if followType == "project" {
			theField = "ProjectFollowers"
		}
	} else if strings.ToLower(field) == "following" {
		if followType == "user" {
			theField = "UserFollowing"
		} else if followType == "project" {
			theField = "ProjectFollowing"
		}
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	return err
}
