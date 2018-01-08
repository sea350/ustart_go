package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendFollow ... appends to either sent or received collegue request arrays within user
//takes in eclient, user ID, the follower ID, and a bool
//true = append to following, false = append to followers
func AppendFollow(eclient *elastic.Client, usrID string, followID string, whichOne bool) error {

	ctx := context.Background()

	FollowLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.Following = append(usr.Following, followID)

		_, err = eclient.Update().
			Index(globals.UserIndex).
			Type(globals.UserType).
			Id(usrID).
			Doc(map[string]interface{}{"Following": usr.Following}).
			Do(ctx)

		return err
	}
	usr.Followers = append(usr.Followers, followID)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Followers": usr.Followers}).
		Do(ctx)

	defer FollowLock.Unlock()
	return err
}
