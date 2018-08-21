package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ReindexProxyNotifications ...
//Reindex proxy notification
func ReindexProxyNotifications(eclient *elastic.Client, docID string, newNotif types.ProxyNotifications) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.ProxyNotifIndex).Do(ctx)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.ProxyNotifIndex).
		Type(globals.ProxyNotifType).
		Id(docID).
		BodyJson(newNotif).
		Do(ctx)

	return nil
}
