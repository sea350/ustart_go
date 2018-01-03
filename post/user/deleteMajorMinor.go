package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteMajorMinor ... appends to either sent or received collegue request arrays within user
//takes in eclient, user ID, the major or minor, an index of the element within the array, and a bool
//true = major, false = minor
func DeleteMajorMinor(eclient *elastic.Client, usrID string, majorMinor string, whichOne bool) error {

	ctx := context.Background()

	procLock.Lock()
	defer procLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		index := -1
		for i := range usr.Majors {
			if usr.Majors[i] == majorMinor {
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

		return err
	}
	index := -1
	for i := range usr.Minors {
		if usr.Minors[i] == majorMinor {
			index = i
		}
	}
	if index < 0 {
		return errors.New("Index not found")
	}
	usr.Minors = append(usr.Minors[:index], usr.Minors[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Minors": usr.Minors}).
		Do(ctx)

	return err
}
