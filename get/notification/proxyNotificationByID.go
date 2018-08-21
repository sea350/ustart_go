package get

import (
	"context"
	"encoding/json"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProxyNotificationByID ...
func ProxyNotificationByID(eclient *elastic.Client, notifID string) (types.ProxyNotifications, error) {

	var notif types.ProxyNotifications

	ctx := context.Background()         //intialize context background
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.ProxyNotifIndex).
						Type(globals.ProxyNotifType).
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
