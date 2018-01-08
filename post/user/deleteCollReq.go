package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteCollReq ...
//  Deletes from sent or received collegue request arrays depending on whichOne
//  True = sent; false = received
func DeleteCollReq(eclient *elastic.Client, usrID string, reqID string, whichOne bool) error {
	ctx := context.Background()

	ColleagueLock.Lock()

	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		index := -1
		for i := range usr.SentCollReq {
			if usr.SentCollReq[i] == reqID {
				index = i
			}
		}
		if index < 0 {
			return errors.New("Index not found")
		}

		usr.SentCollReq = append(usr.SentCollReq[:index], usr.SentCollReq[index+1:]...)

		_, err = eclient.Update().
			Index(globals.UserIndex).
			Type(globals.UserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentCollReq": usr.SentCollReq}).
			Do(ctx)

		return err
	}

	index := -1
	for i := range usr.ReceivedCollReq {
		if usr.SentCollReq[i] == reqID {
			index = i
		}
	}
	if index < 0 {
		return errors.New("Index not found")
	}
	usr.ReceivedCollReq = append(usr.ReceivedCollReq[:index], usr.ReceivedCollReq[index+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedCollReq": usr.ReceivedCollReq}).
		Do(ctx)

	defer ColleagueLock.Unlock()
	return err
}
