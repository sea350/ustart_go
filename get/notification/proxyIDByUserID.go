package get

import (
	"context"
	"errors"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProxyIDByUserID ...
func ProxyIDByUserID(eclient *elastic.Client, userID string) (string, error) {

	ctx := context.Background()

	termQuery := elastic.NewTermQuery("DocID", strings.ToLower(userID))
	searchResult, err := eclient.Search().
		Index(globals.ProxyNotifIndex).
		Type(globals.ProxyNotifType).
		Query(termQuery).
		Do(ctx)

	var proxyID string
	if err != nil {
		return proxyID, err
	}

	if searchResult.TotalHits() == 0 {
		exists, _ := eclient.IndexExists(globals.ProxyNotifIndex).Do(ctx)
		if !exists {
			_, err := eclient.CreateIndex(globals.ProxyNotifIndex).Do(ctx)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				return proxyID, err
			}
		}
		var settings types.NotificationSettings
		settings.Default()
		idx, err := eclient.Index().
			Index(globals.ProxyNotifIndex).
			BodyJson(types.ProxyNotifications{DocID: userID, Settings: settings}).
			Do(ctx)

		return idx.Id, err
	}
	if searchResult.TotalHits() > 1 {
		return proxyID, errors.New("multiple proxies found")
	}

	for _, element := range searchResult.Hits.Hits {
		proxyID = element.Id
		break
	}

	return proxyID, err

}
