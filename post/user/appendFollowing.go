package post

import (
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendFollowing ... appends to user's following array
//takes in eclient, user ID, the followee ID
func AppendFollowing(eclient *elastic.Client, usrID string, followID string) error {

	FollowLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Following = append(usr.Following, followID)

	err = UpdateUser(eclient, usrID, "Following", usr.Following)

	defer FollowLock.Unlock()
	return err
}
