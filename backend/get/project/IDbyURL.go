package get

import (
	"context"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectIDByURL ...
func ProjectIDByURL(eclient *elastic.Client, projectURL string) (string, error) {
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT URL)
	//RETURNS A types.Project AND AN error
	ctx := context.Background()
	//Need to lowercase to query
	termQuery := elastic.NewTermQuery("URLName", strings.ToLower(projectURL))
	searchResult, err := eclient.Search().
		Index(globals.ProjectIndex).
		Query(termQuery).
		Do(ctx)

	var result string

	if searchResult.Hits.TotalHits > 1 {
		return "", errors.New("Too many results")
	} else if searchResult.Hits.TotalHits < 1 {
		return "", errors.New("No results")
	}
	for _, element := range searchResult.Hits.Hits {

		result = element.Id
		break
	}

	return result, err

}
