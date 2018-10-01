package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteSentEventReq ... whichOne: true = sent
//whichOne: false = received
func DeleteSentEventReq(eclient *elastic.Client, usrID string, eventID string) error {
	ctx := context.Background()

	EventLock.Lock()
	defer EventLock.Unlock()

	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := -1
	for i := range usr.SentEventReq {
		if usr.SentEventReq[i] == eventID {
			index = i
			break
		}
	}

	if index < 0 {
		return errors.New("index does not exist")
	}
	//end of temp solution

	usr.SentEventReq = append(usr.SentEventReq[:index], usr.SentEventReq[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"SentEventReq": usr.SentEventReq}).
		Do(ctx)

	return err

}
