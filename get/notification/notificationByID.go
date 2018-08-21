package get

import (
	"context"
	"encoding/json"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NotificationByID ...
func NotificationByID(eclient *elastic.Client, notifID string) (types.Notification, error) {

	var notif types.Notification

	ctx := context.Background()         //intialize context background
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.NotificationIndex).
						Type(globals.NotificationType).
						Id(notifID).
						Do(ctx)
	if err != nil {
		return notif, err
	}

	Err := json.Unmarshal(*searchResult.Source, &notif)
	if Err != nil {
		return notif, Err
	} //forward error

	return notif, Err

}
