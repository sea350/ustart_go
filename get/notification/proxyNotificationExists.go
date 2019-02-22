package get

import (
	"context"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//ProxyNotificationExists ... returns proxy id and struct
func ProxyNotificationExists(eclient *elastic.Client, userID string) (bool, error) {

	ctx := context.Background()
	// var proxy types.ProxyNotifications

	exists, err := eclient.IndexExists(globals.ProxyNotifIndex).Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	if !exists {
		_, err := eclient.CreateIndex(globals.ProxyNotifIndex).Do(ctx)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return false, err
		}
	}
	termQuery := elastic.NewTermQuery("DocID", strings.ToLower(userID))
	searchResult, err := eclient.Search().
		Index(globals.ProxyNotifIndex).
		Type(globals.ProxyNotifType).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return false, err
	}

	return searchResult.TotalHits() == 0, err

}
