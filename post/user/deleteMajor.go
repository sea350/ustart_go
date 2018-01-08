package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteMajor ... deletes form majpr array
//takes in eclient, user ID, the major or minor, an index of the element within the array
func DeleteMajor(eclient *elastic.Client, usrID string, major string) error {
	ctx := context.Background()

	procLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := -1
	for i := range usr.Majors {
		if usr.Majors[i] == major {
			index = i
		}
	}
	if index < 0 {
		return errors.New("Index not found")
	}
	usr.Majors = append(usr.Majors[:index], usr.Majors[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Majors": usr.Majors}).
		Do(ctx)

	defer procLock.Unlock()
	return err

}
