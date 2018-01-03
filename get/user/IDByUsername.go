package get

import (
	"context"
	"errors"

	elastic "gopkg.in/olivere/elastic.v5"
)

func GetIDByUsername(eclient *elastic.Client, username string) (string, error) {
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("Username", username)
	searchResult, err := eclient.Search().
		Index(USER_INDEX).
		Query(termQuery).
		Do(ctx)

	var result string //save id to a result variable
	if searchResult.TotalHits() > 1 {
		return result, errors.New("More than one user found")
	}

	for _, element := range searchResult.Hits.Hits { //interate through hits, get the element id
		result = element.Id

	}

	return result, err //return id
}
