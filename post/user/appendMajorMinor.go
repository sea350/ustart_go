package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendMajorMinor ...
func AppendMajorMinor(eclient *elastic.Client, usrID string, majorMinor string, whichOne bool) error {
	//appends to either sent or received collegue request arrays within user
	//takes in eclient, user ID, the major or minor, and a bool
	//true = major, false = minor
	ctx := context.Background()

	procLock.Lock()
	defer procLock.Unlock()
	//
	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.Majors = append(usr.Majors, majorMinor)

		_, err = eclient.Update().
			Index(globals.UserIndex).
			Type(globals.UserType).
			Id(usrID).
			Doc(map[string]interface{}{"Majors": usr.Majors}).
			Do(ctx)

		return err
	}
	usr.Minors = append(usr.Minors, majorMinor)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Minors": usr.Minors}).
		Do(ctx)

	return err

}
