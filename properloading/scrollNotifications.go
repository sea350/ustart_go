package properloading

import (
	"context"
	"encoding/json"
	"errors"
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

	var mapResults map[string]types.Notification

	scroll := eclient.Scroll().
		Index(globals.NotificationIndex).
		Query(notifQuery).
		Sort("TimeStamp", false).
		Size(5)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)
	if err == io.EOF {
		return "", mapResults, 0, err //we might need special treatment for EOF error
	}
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return "", mapResults, 0, err
	}

	var notif types.Notification

	for _, hit := range res.Hits.Hits {
		// fmt.Println(hit.Id)
		err = json.Unmarshal(*hit.Source, &notif)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("ISSUE WITH CONVERT FUNCTION")
			continue
		}

		mapResults[hit.Id] = notif

		if err == io.EOF {
			return res.ScrollId, mapResults, int(res.Hits.TotalHits), errors.New("Out of bounds")

		}

	}

	return res.ScrollId, mapResults, int(res.Hits.TotalHits), err
}
