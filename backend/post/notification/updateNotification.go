package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/notification"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateNotification ... UPDATES A SINGLE FEILD IN AN EXISTING ES DOC
//Returns an error
func UpdateNotification(eclient *elastic.Client, notificationID string, field string, newContent interface{}) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.NotificationIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.NotificationByID(eclient, notificationID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.NotificationIndex).
		Type(globals.NotificationType).
		Id(notificationID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	return err
}
