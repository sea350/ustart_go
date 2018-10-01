package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/chat"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateProxyMsg  ...
//Updates proxy messages
func UpdateProxyMsg(eclient *elastic.Client, msgID string, field string, newContent interface{}) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.ProxyMsgIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.ProxyMsgByID(eclient, msgID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.ProxyMsgIndex).
		Type(globals.ProxyMsgType).
		Id(msgID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	return nil
}
