package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//EventIDByURL ... pulls an event from ES, returns event ID
func EventIDByURL(eclient *elastic.Client, eventURL string) (string, error) {
	ctx := context.Background()

	termQuery := elastic.NewTermQuery("URLName", strings.ToLower(eventURL))
	searchResult, err := eclient.Search().
		Index(globals.EventIndex).
		Query(termQuery).
		Do(ctx)

	var result string

	for _, element := range searchResult.Hits.Hits {
		result = element.Id
		break
	}

	return result, err
}
