package get

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProxyNotificationByUserID ...
func ProxyNotificationByUserID(eclient *elastic.Client, userID string) (types.ProxyNotifications, error) {

	ctx := context.Background()

	termQuery := elastic.NewTermQuery("DocID", strings.ToLower(userID))
	searchResult, err := eclient.Search().
		Index(globals.ProxyNotifIndex).
		Query(termQuery).
		Do(ctx)

	var proxy types.ProxyNotifications
	if err != nil {
		return proxy, err
	}

	if searchResult.TotalHits() == 0 {
		return proxy, errors.New("No results, proxy ID does not exist")
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
