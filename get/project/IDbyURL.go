package get

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectIDByURL ...
func ProjectIDByURL(eclient *elastic.Client, projectURL string) (string, error) {
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT URL)
	//RETURNS A types.Project AND AN error
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("URL", projectURL)
	searchResult, err := eclient.Search().
		Index(globals.ProjectIndex).
		Query(termQuery).
		Do(ctx)

	var result string

	for _, element := range searchResult.Hits.Hits {

		result = element.Id
		break
	}

	return result, err

}
