package get

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/backend/globals"
	"github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProxyNotificationByUserID ... returns proxy id and struct
func ProxyNotificationByUserID(eclient *elastic.Client, userID string) (string, types.ProxyNotifications, error) {

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
			return ``, proxy, err
		}
	}
	termQuery := elastic.NewTermQuery("DocID", strings.ToLower(userID))
	searchResult, err := eclient.Search().
		Index(globals.ProxyNotifIndex).
		Type(globals.ProxyNotifType).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return ``, proxy, err
	}

	if searchResult.TotalHits() == 0 {
		proxy.Settings.Default()
		proxy.DocID = userID

		res, err := eclient.Index().
			Index(globals.ProxyNotifIndex).
			Type(globals.ProxyNotifType).
			BodyJson(proxy).
			Do(ctx)

		if err != nil {
			return res.Id, types.ProxyNotifications{}, err
		}
		return res.Id, proxy, nil
	}
	if searchResult.TotalHits() > 1 {
		return "", proxy, errors.New("multiple proxies found")
	}

	var id string
	for _, element := range searchResult.Hits.Hits {
		Err := json.Unmarshal(*element.Source, &proxy)
		id = element.Id
		if Err != nil {
			return element.Id, proxy, Err
		}
	}

	return id, proxy, err

}
