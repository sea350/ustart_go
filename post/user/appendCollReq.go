package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendCollReq ...
//  Appends to either sent or received collegue, based on whichOne
//  True = sent; False = received.
func AppendCollReq(eclient *elastic.Client, usrID string, collegueID string, whichOne bool) error {

	ctx := context.Background()

	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {

		usr.SentCollReq = append(usr.SentCollReq, collegueID)

		_, err = eclient.Update().
			Index(globals.UserIndex).
			Type(globals.UserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentCollReq": usr.SentCollReq}).
			Do(ctx)

		return err
	}

	usr.ReceivedCollReq = append(usr.ReceivedCollReq, collegueID)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedCollReq": usr.ReceivedCollReq}).
		Do(ctx)

	return err
}
