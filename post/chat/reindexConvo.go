package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ReindexConversation ...
//Reindexes conversation
func ReindexConversation(eclient *elastic.Client, newConvo types.Conversation, convoID string) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.ConvoIndex).Do(ctx)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.ConvoIndex).
		Type(globals.ConvoType).
		Id(convoID).
		BodyJson(newConvo).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
