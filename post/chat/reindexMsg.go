package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ReindexMsg ...
//Reindex messages
func ReindexMsg(eclient *elastic.Client, newMsg types.Message, msgID string) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.MsgIndex).Do(ctx)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.MsgIndex).
		Type(globals.MsgType).
		Id(msgID).
		BodyJson(newMsg).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
