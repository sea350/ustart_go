package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteProjReq ... whichOne: true = sent
//whichOne: false = received
func DeleteProjReq(eclient *elastic.Client, usrID string, projID string, whichOne bool) error {
	ctx := context.Background()

	projectLock.Lock()
	defer projectLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		//universal.FindIndex(usr.SentProjReq, projID)
		//temp solution
		index := -1
		for i := range usr.SentProjReq {
			if usr.SentProjReq[i] == projID {
				index = i
				break
			}
		}

		if index < 0 {
			return errors.New("index does not exist")
		}
		//end of temp solution

		usr.SentProjReq = append(usr.SentProjReq[:index], usr.SentProjReq[index+1:]...)

		_, err = eclient.Update().
			Index(globals.UserIndex).
			Type(globals.UserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentProjReq": usr.SentProjReq}).
			Do(ctx)

		return err
	}
	//universal.FindIndex(usr.ReceivedProjReq, projID)
	//temp solution
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

	return err
}
