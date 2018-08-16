package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventByURL ... queries ES to get the event by ID
func EventByURL(eclient *elastic.Client, eventURL string) (types.Events, error) {
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("URLName", strings.ToLower(eventURL))
	searchResult, err := eclient.Search().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Query(termQuery).
		Do(ctx)

	var result string
	var evnt types.Events
	for _, element := range searchResult.Hits.Hits {
		result = element.Id
		break
	}

	evnt, err = EventByID(eclient, result)
	return evnt, err
}
