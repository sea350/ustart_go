package get

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IDByUsername ...
func IDByUsername(eclient *elastic.Client, username string) (string, error) {
	ctx := context.Background()
	termQuery := elastic.NewMatchQuery("Username", username)
	searchResult, err := eclient.Search().
		Index(globals.UserIndex).
		Query(termQuery).
		Do(ctx)

	var result string //save id to a result variable
	if searchResult.TotalHits() > 1 {
		return result, errors.New("More than one user found")
	}
	if searchResult.TotalHits() == 0 {
		return result, errors.New("No users found")
	}

	for _, element := range searchResult.Hits.Hits { //interate through hits, get the element id
		result = element.Id

	}

	return result, err //return id
}
