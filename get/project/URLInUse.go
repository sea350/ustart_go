package get

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//URLInUse ...
func URLInUse(eclient *elastic.Client, projectURL string) (bool, error) {
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT URL)
	//RETURNS A types.Project AND AN error
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("URL", projectURL)
	searchResult, err := eclient.Search().
		Index(globals.ProjectIndex).
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
