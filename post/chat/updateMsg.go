package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/chat"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//UpdateMsg ...
//Updates messages
func UpdateMsg(eclient *elastic.Client, chatID string, field string, newContent interface{}) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.MsgIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.MsgByID(eclient, chatID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.MsgIndex).
		Type(globals.MsgType).
		Id(chatID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
