package post

import (
	"context"
	"errors"
	"fmt"

	get "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteFollow ... whichOne: true = following
//whichOne: false = followers
//followID does nothing
func DeleteFollow(eclient *elastic.Client, usrID string, followID string, whichOne bool) error {

	ctx := context.Background()

	FollowLock.Lock()
	defer FollowLock.Unlock()
	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		index := -1
		for i := range usr.Following {
			if usr.Following[i] == followID {
				fmt.Println("current ID:", usr.Following[i])
				index = i
			}
		}
		if index < 0 {
			return errors.New("Index not found")
		}
		usr.Following = append(usr.Following[:index], usr.Following[index+1:]...)

		_, err = eclient.Update().
			Index(globals.UserIndex).
			Type(globals.UserType).
			Id(usrID).
			Doc(map[string]interface{}{"Following": usr.Following}).
			Do(ctx)

		return err

	}
	index := -1
	for i := range usr.Followers {
		if usr.Followers[i] == followID {
			index = i
		}
	}
	if index < 0 {
		return errors.New("Index not found")
	}
	usr.Followers = append(usr.Followers[:index], usr.Followers[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Followers": usr.Followers}).
		Do(ctx)

	return err
}
