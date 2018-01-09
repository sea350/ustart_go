package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteReceivedProjReq ...
func DeleteReceivedProjReq(eclient *elastic.Client, usrID string, projID string) error {
	ctx := context.Background()

	ProjectLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := 0
	for i := range usr.ReceivedProjReq {
		if usr.ReceivedProjReq[i] == projID {
			index = i
			break
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	//end of temp solution
	usr.ReceivedProjReq = append(usr.ReceivedProjReq[:index], usr.ReceivedProjReq[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedProjReq": usr.ReceivedProjReq}).
		Do(ctx)

	defer ProjectLock.Unlock()
	return err
}
