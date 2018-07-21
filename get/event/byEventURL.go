package get

import (
	"context"
	"fmt"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventByURL ... queries ES to get the event by URL
func EventByURL(eclient *elastic.Client, eventURL string) (types.Events, error) {
	fmt.Println("1")
	ctx := context.Background()
	fmt.Println("2")
	termQuery := elastic.NewTermQuery("URLName", strings.ToLower(eventURL))
	fmt.Println("3")
	searchResult, err := eclient.Search().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Query(termQuery).
		Do(ctx)
	fmt.Println("4")
	var result string
	var evnt types.Events
	for _, element := range searchResult.Hits.Hits {
		fmt.Println("5")
		result = element.Id
		break
	}

	evnt, err = EventByID(eclient, result)

	return evnt, err
}
