package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteReceivedEventReq ...
func DeleteReceivedEventReq(eclient *elastic.Client, usrID string, eventID string) error {
	ctx := context.Background()

	EventLock.Lock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := 0
	for i := range usr.ReceivedEventReq {
		if usr.ReceivedEventReq[i] == eventID {
			index = i
			break
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	//end of temp solution
	usr.ReceivedEventReq = append(usr.ReceivedEventReq[:index], usr.ReceivedEventReq[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedEventReq": usr.ReceivedEventReq}).
		Do(ctx)

	defer EventLock.Unlock()
	return err
}
