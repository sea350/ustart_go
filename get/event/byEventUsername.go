package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NOT NEEDED! WILL GET BACK TO THIS LATER IF WE DO NEED IT!

//EventByUsername ...
func EventByUsername(eclient *elastic.Client, username string) (types.Events, error) {
	ctx := context.Background()

	termQuery := elastic.NewTermQuery("Username", strings.ToLower(username))
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
