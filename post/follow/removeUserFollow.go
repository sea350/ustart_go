package post

import (
	"context"
	"errors"
	"strings"

	getFollow "github.com/sea350/ustart_go/get/follow"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//RemoveUserFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func RemoveUserFollow(eclient *elastic.Client, userID string, field string, deleteKey string, followType string) error {

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
		if followType == "user" {
			FollowLock.Lock()
			defer FollowLock.Unlock()
			if len(foll.UserFollowers) == 0 {
				return errors.New("No followers to remove")
			}
			delete(foll.UserFollowers, deleteKey)

		} else if followType == "project" {
			FollowLock.Lock()
			defer FollowLock.Unlock()
			if len(foll.ProjectFollowers) == 0 {
				return errors.New("No followers to remove")
			}
			delete(foll.ProjectFollowers, deleteKey)
		}
	case "following":
		FollowLock.Lock()
		defer FollowLock.Unlock()
		if followType == "user" {
			if len(foll.UserFollowing) == 0 {
				return errors.New("Nothing to remove from following")
			}
			delete(foll.UserFollowing, deleteKey)
		} else if followType == "project" {
			if len(foll.ProjectFollowing) == 0 {
				return errors.New("Nothing to remove from following")
			}
			delete(foll.ProjectFollowing, deleteKey)
		}

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
