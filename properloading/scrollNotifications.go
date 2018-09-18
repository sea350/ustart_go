package properloading

import (
	"context"
	"encoding/json"
	"io"
	"log"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ScrollNotifications ...
//Scrolls through docs being loaded
func ScrollNotifications(eclient *elastic.Client, docID string, scrollID string) (string, map[string]types.Notification, int, error) {

	ctx := context.Background()

	//set up user query
	notifQuery := elastic.NewBoolQuery()
	notifQuery = notifQuery.Must(elastic.NewTermQuery("DocID", docID))

	notifQuery = notifQuery.Must(elastic.NewTermQuery("Invisible", false))

	//yeah....

	var mapResults = make(map[string]types.Notification)

	scroll := eclient.Scroll().
		Index(globals.NotificationIndex).
		Query(notifQuery).
		Sort("Timestamp", false).
		Size(5)

	if len(scrollID) > 0 {
		scroll = scroll.ScrollId(scrollID)
	}

	res, scrollErr := scroll.Do(ctx)
	// if err == io.EOF {
	// 	return "", mapResults, 0, err //we might need special treatment for EOF error
	// }
	if scrollErr != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(scrollErr)
		return "", mapResults, 0, scrollErr
	}

	var notif types.Notification

	for _, hit := range res.Hits.Hits {
		// fmt.Println(hit.Id)
		if scrollErr == io.EOF {
			continue
		}
		err := json.Unmarshal(*hit.Source, &notif)
		if err != nil && err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)

			continue
		}

		mapResults[hit.Id] = notif

	}

	return res.ScrollId, mapResults, int(res.Hits.TotalHits), nil
}
