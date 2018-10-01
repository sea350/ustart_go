package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteColleague ... deletes from collegue array within user
//takes in eclient, user ID, and collegue ID
func DeleteColleague(eclient *elastic.Client, usrID string, deleteID string) error {
	ctx := context.Background()

	ColleagueLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	//idx, err := universal.FindIndex(usr.Colleagues, deleteID) UNIVERSAL PKG
	//temp for-loop:

	index := -1
	for i := range usr.Colleagues {
		if usr.Colleagues[i] == deleteID {
			index = i
		}
	}

	if index < 0 {
		return errors.New("Index non-existent")
	}
	//temp solution stops here

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Colleagues = append(usr.Colleagues[:index], usr.Colleagues[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Colleagues": usr.Colleagues}).
		Do(ctx)

	defer ColleagueLock.Unlock()
	return err

}
