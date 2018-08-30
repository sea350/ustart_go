package get

import (
	"context"
	"encoding/json"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NotificationsByUserID ...
func NotificationsByUserID(eclient *elastic.Client, userID string) ([]types.Notification, error) {

	ctx := context.Background()

	termQuery := elastic.NewTermQuery("DocID", strings.ToLower(userID))
	searchResult, err := eclient.Search().
		Index(globals.NotificationIndex).
		Type(globals.NotificationType).
		Query(termQuery).
		Do(ctx)

	var notifs []types.Notification
	if err != nil {
		return notifs, err
	}

	for _, element := range searchResult.Hits.Hits {
		var temp types.Notification
		err := json.Unmarshal(*element.Source, &temp)
		if err != nil {
			continue
		}
		notifs = append(notifs, temp)
	}

	return notifs, err

}
