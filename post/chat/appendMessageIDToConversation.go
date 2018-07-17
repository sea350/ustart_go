package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/chat"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendMessageIDToConversation ...
//Appends with its own special "lock" used for concurrency control
func AppendMessageIDToConversation(eclient *elastic.Client, convoID string, newMessageID string) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.ConvoIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	AppendMessageLock.Lock()
	defer AppendMessageLock.Unlock()

	convo, err := get.ConvoByID(eclient, convoID)
	if err != nil {
		return err
	}

	convo.MessageIDArchive = append(convo.MessageIDArchive, newMessageID)
	//add any cache control here if necessary

	_, err = eclient.Update().
		Index(globals.ConvoIndex).
		Type(globals.ConvoType).
		Id(convoID).
		Doc(map[string]interface{}{"MessageIDArchive": convo.MessageIDArchive}).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
