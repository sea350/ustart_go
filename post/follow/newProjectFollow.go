package post

import (
	"context"
	"errors"
	"strings"

	getFollow "github.com/sea350/ustart_go/get/follow"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NewProjectFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func NewProjectFollow(eclient *elastic.Client, userID string, field string, newKey string, isBell bool) error {

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
		FollowerLock.Lock()
		defer FollowerLock.Unlock()
		foll.ProjectFollowers[newKey] = isBell
		followMap = foll.ProjectFollowers

	case "following":
		FollowingLock.Lock()
		defer FollowingLock.Unlock()
		foll.ProjectFollowing[newKey] = isBell
		followMap = foll.ProjectFollowing
	default:
		return errors.New("Invalid field")
	}
	_, err = eclient.Update().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		Id(follID).
		Doc(map[string]interface{}{field: followMap}). //field = Followers or Following, newContent =
		Do(ctx)

	return err
}
