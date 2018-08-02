package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ReindexProxyMsg ...
//Reindex messages
func ReindexProxyMsg(eclient *elastic.Client, msgID string, newMsg types.ProxyMessages) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.ProxyMsgIndex).Do(ctx)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.ProxyMsgIndex).
		Type(globals.ProxyMsgType).
		Id(msgID).
		BodyJson(newMsg).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
