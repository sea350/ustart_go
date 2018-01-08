package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteMinor ... deletes from minor array
//takes in eclient, user ID, the minor, an index of the element within the array
func DeleteMinor(eclient *elastic.Client, usrID string, minor string) error {

	ctx := context.Background()

	procLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := -1
	for i := range usr.Minors {
		if usr.Majors[i] == minor {
			index = i
		}
	}
	if index < 0 {
		return errors.New("Index not found")
	}
	usr.Majors = append(usr.Minors[:index], usr.Minors[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Minors": usr.Minors}).
		Do(ctx)

	defer procLock.Unlock()
	return err

}
