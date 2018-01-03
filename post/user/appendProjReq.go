package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendProjReq ... appends to either sent or received project request arrays within user
//takes in eclient, user ID, the project ID, and a bool
//true = append to following, false = append to followers
func AppendProjReq(eclient *elastic.Client, usrID string, projID string, whichOne bool) error {
	ctx := context.Background()

	projectLock.Lock()
	defer procLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.SentProjReq = append(usr.SentProjReq, projID)

		_, err = eclient.Update().
			Index(globals.UserIndex).
			Type(globals.UserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentProjReq": usr.SentProjReq}).
			Do(ctx)

		return err
	}
	usr.ReceivedProjReq = append(usr.ReceivedProjReq, projID)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedProjReq": usr.ReceivedProjReq}).
		Do(ctx)

	return err
}
