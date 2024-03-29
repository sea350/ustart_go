package get

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//WidgetByLink ...
func WidgetByLink(eclient *elastic.Client, link string) (types.Widget, error) {

	ctx := context.Background()

	termQuery := elastic.NewTermQuery("Link", link)
	searchResult, err := eclient.Search().
		Index(globals.WidgetIndex).
		Query(termQuery).
		Do(ctx)

	var widget types.Widget
	if err != nil {
		return widget, err
	}
	var result string

	for _, element := range searchResult.Hits.Hits {

		result = element.Id
		break
	}

	widget, err = WidgetByID(eclient, result)

	return widget, err

}
