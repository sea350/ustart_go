package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendMajor ...
func AppendMajor(eclient *elastic.Client, usrID string, major string) error {
	//appends to either sent or received collegue request arrays within user

	ctx := context.Background()

	procLock.Lock()

	//
	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Majors = append(usr.Majors, major)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Majors": usr.Majors}).
		Do(ctx)

	defer procLock.Unlock()
	return err

}
