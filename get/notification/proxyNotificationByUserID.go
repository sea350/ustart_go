package get

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProxyNotificationByUserID ...
func ProxyNotificationByUserID(eclient *elastic.Client, userID string) (types.ProxyNotifications, error) {

	ctx := context.Background()
	var proxy types.ProxyNotifications

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
			return proxy, err
		}
	}
	termQuery := elastic.NewTermQuery("DocID", strings.ToLower(userID))
	searchResult, err := eclient.Search().
		Index(globals.ProxyNotifIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return proxy, err
	}

	if searchResult.TotalHits() == 0 {
		proxy.Settings.Default()
		proxy.DocID = userID

		_, err := eclient.Index().
			Index(globals.ProxyNotifIndex).
			BodyJson(proxy).
			Do(ctx)

		if err != nil {
			return types.ProxyNotifications{}, err
		} else {
			return proxy, nil
		}
	}
	if searchResult.TotalHits() > 1 {
		return proxy, errors.New("multiple proxies found")
	}

	for _, element := range searchResult.Hits.Hits {
		Err := json.Unmarshal(*element.Source, &proxy)
		if Err != nil {
			return proxy, Err
		}
	}

	return proxy, err

}
