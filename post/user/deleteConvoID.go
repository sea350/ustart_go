package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteConvoID ...
func DeleteConvoID(eclient *elastic.Client, usrID string, convoID string, idx int) error {
	ctx := context.Background()
	usr, err := get.UserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	usr.ConversationIDs = append(usr.ConversationIDs[:idx], usr.ConversationIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"ConversationIDs": usr.ConversationIDs}).
		Do(ctx)

	return err

}
