package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/chat"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateConvo ...
//Updates conversation
func UpdateConvo(eclient *elastic.Client, convoID string, field string, newContent interface{}) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.ConvoIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.ConvoByID(eclient, chatID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.ConvoIndex).
		Type(globals.ConvoType).
		Id(convoID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
