package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendMinor ...
func AppendMinor(eclient *elastic.Client, usrID string, minor string) error {
	//appends to either sent or received collegue request arrays within user
	//takes in eclient, user ID, the major or minor, and a bool
	//true = major, false = minor
	ctx := context.Background()

	procLock.Lock()

	//
	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Minors = append(usr.Minors, minor)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Minors": usr.Minors}).
		Do(ctx)

	defer procLock.Unlock()

	return err

}
