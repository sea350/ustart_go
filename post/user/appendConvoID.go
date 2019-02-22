package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//AppendConvoID ...
func AppendConvoID(eclient *elastic.Client, usrID string, convoID string) error {
	ctx := context.Background()
	usr, err := get.UserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.ConversationIDs = append(usr.ConversationIDs, convoID)

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(usrID).
		Doc(map[string]interface{}{"ConversationIDs": usr.ConversationIDs}).
		Do(ctx)

	return err

}
