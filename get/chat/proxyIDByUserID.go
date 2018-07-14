package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProxyIDByUserID ...
func ProxyIDByUserID(eclient *elastic.Client, userID string) (string, error) {

	ctx := context.Background()

	termQuery := elastic.NewTermQuery("DocID", strings.ToLower(userID))
	searchResult, err := eclient.Search().
		Index(globals.ProxyMsgIndex).
		Query(termQuery).
		Do(ctx)

	var proxyID string
	if err != nil {
		return proxyID, err
	}

	for _, element := range searchResult.Hits.Hits {

		proxyID = element.Id
		break
	}

	return proxyID, err

}
