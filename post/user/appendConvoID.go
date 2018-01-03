package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendConvoID ...
func AppendConvoID(eclient *elastic.Client, usrID string, convoID string) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

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
