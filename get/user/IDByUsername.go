package get

import (
	"context"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//IDByUsername ...
func IDByUsername(eclient *elastic.Client, username string) (string, error) {
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("Username", strings.ToLower(username))
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
