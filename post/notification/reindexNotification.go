package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ReindexNotification ...
//Reindex notification
func ReindexNotification(eclient *elastic.Client, newNotif types.Notification, docID string) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.NotificationIndex).Do(ctx)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.NotificationIndex).
		Type(globals.NotificationType).
		Id(docID).
		BodyJson(newNotif).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
