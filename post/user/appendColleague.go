package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendColleague ... appends to collegue array within user
//takes in eclient, user ID, and collegue ID
func AppendColleague(eclient *elastic.Client, usrID string, colleagueID string) error {

	ctx := context.Background()

	ColleagueLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Colleagues = append(usr.Colleagues, colleagueID)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"Colleagues": usr.Colleagues}).
		Do(ctx)

	defer ColleagueLock.Unlock()
	return err

}
