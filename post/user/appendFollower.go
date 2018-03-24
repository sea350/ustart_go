package post

import (
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendFollower ...
//takes in eclient, user ID, the follower ID
func AppendFollower(eclient *elastic.Client, usrID string, followID string) error {

	FollowLock.Lock()
	defer FollowLock.Unlock()
	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Followers = append(usr.Followers, followID)

	err = UpdateUser(eclient, usrID, "Followers", usr.Followers)

	return err
}
