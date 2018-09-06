package post

import (
	"context"
	"errors"
	"log"

	get "github.com/sea350/ustart_go/get/notification"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateProxyNotifications  ...
//Updates proxy notifications
func UpdateProxyNotifications(eclient *elastic.Client, msgID string, field string, newContent interface{}) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.ProxyNotifIndex).Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.ProxyNotificationByID(eclient, msgID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	_, err = eclient.Update().
		Index(globals.ProxyNotifIndex).
		Type(globals.ProxyNotifType).
		Id(msgID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	return nil
}
