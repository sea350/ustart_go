package post

import (
	"context"
	"log"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//IndexNotification ...
//Indexes a new notification, also adds to aphiliated proxies and will merge notifications when necessary
//returns the new ID and error
func IndexNotification(eclient *elastic.Client, newNotif types.Notification) (string, error) {
	ctx := context.Background()
	var notifID string
	exists, err := eclient.IndexExists(globals.NotificationIndex).Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	if !exists {
		_, err := eclient.CreateIndex(globals.NotificationIndex).Do(ctx)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}

	}
	idx, Err := eclient.Index().
		Index(globals.NotificationIndex).
		Type(globals.NotificationType).
		BodyJson(newNotif).
		Do(ctx)

	if Err != nil {
		return notifID, Err
	}
	notifID = idx.Id

	err = AppendToProxyNotificationsByUserID(eclient, newNotif.DocID, notifID)

	return notifID, err
}
