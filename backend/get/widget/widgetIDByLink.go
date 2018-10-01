package get

import (
	"context"

	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//WidgetIDByLink ...
func WidgetIDByLink(eclient *elastic.Client, link string) (string, error) {

	ctx := context.Background()

	//username:= EmailToUsername(email) //for username query
	termQuery := elastic.NewTermQuery("Link", link)
	searchResult, err := eclient.Search().
		Index(globals.WidgetIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return "", err
	}
	var result string

	for _, element := range searchResult.Hits.Hits {

		result = element.Id
		break
	}

	return result, err

}
