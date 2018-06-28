package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventByEmail ...
func EventByEmail(eclient *elastic.Client, email string) (types.Events, error) {
	ctx := context.Background()

	termQuery := elastic.NewTermQuery("Email", strings.ToLower(email))
	searchResult, err := eclient.Search().
		Index(globals.EventIndex).
		Query(termQuery).
		Do(ctx)

	var evnt types.Events
	if err != nil {
		return evnt, err
	}

	var result string
	for _, element := range searchResult.Hits.Hits {
		result = element.Id
		break
	}

	evnt, err = EventByID(eclient, result)

	return evnt, err

}
