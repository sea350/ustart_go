package post

import (
	"context"
	"errors"

	getFollow "github.com/sea350/ustart_go/get/follow"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//NewUserBell ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func NewUserBell(eclient *elastic.Client, userID string, newKey string, isBell bool) error {

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

	FollowLock.Lock()
	defer FollowLock.Unlock()
	foll.UserFollowers[newKey] = isBell
	followMap = foll.UserFollowers
	//modify user bell map if bell follower
	if isBell {
		foll.UserBell[newKey] = isBell
		bellMap = foll.UserBell
	}

	newFollow := eclient.Update().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		Id(follID).
		Doc(map[string]interface{}{"UserBell": followMap}) //field = Followers or Following, newContent =

	//only executes when there is a new bell follower
	if isBell {
		newFollow.Doc(map[string]interface{}{"UserBell": bellMap})
	}
	_, err = newFollow.Do(ctx)
	return err
}
