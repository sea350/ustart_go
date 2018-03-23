package post

import (
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendFollow ... appends to either sent or received collegue request arrays within user
//takes in eclient, user ID, the follower ID, and a bool
//true = append to following, false = append to followers
func AppendFollower(eclient *elastic.Client, usrID string, followID string) error {

	FollowLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Followers = append(usr.Followers, followID)

	err = UpdateUser(eclient, usrID, "Followers", usr.Followers)

	defer FollowLock.Unlock()
	return err
}
