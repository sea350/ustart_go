package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//UnseenNotificationIDsByUserID ...
func UnseenNotificationIDsByUserID(eclient *elastic.Client, userID string) ([]string, error) {

	ctx := context.Background()

	termQuery := elastic.NewBoolQuery()
	termQuery = termQuery.Must(elastic.NewTermQuery("DocID", strings.ToLower(userID)))
	termQuery = termQuery.Must(elastic.NewTermQuery("Seen", false))
	searchResult, err := eclient.Search().
		Index(globals.NotificationIndex).
		Type(globals.NotificationType).
		Query(termQuery).
		Do(ctx)

	var notifs []string
	if err != nil {
		return notifs, err
	}

	if searchResult.TotalHits() == 0 {
		return notifs, err
	}

	for _, element := range searchResult.Hits.Hits {
		notifs = append(notifs, element.Id)
	}

	return notifs, err

}
