package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventURLInUse ...
func EventURLInUse(eclient *elastic.Client, eventURL string) (bool, error) {
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("URLName", strings.ToLower(eventURL))
	searchResult, err := eclient.Search().
		Index(globals.EventIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return true, err
	}

	if searchResult.Hits.TotalHits > 0 {
		return true, nil
	}

	return false, nil
}
