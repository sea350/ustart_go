package search

import (
	"context"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "github.com/olivere/elastic"
)

//Skills ... Attempt at general skills search
func Skills(eclient *elastic.Client, searchTerm string) ([]types.FloatingHead, error) {

	ctx := context.Background()
	var results []types.FloatingHead
	query := elastic.NewBoolQuery()

	stringArray := strings.Split(searchTerm, ` `)
	for _, element := range stringArray {
		query = query.Should(elastic.NewWildcardQuery("Tags", strings.ToLower(element)))
	}

	searchResults, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(query).
		Do(ctx)

	if err != nil {
		return results, err
	}

	for _, element := range searchResults.Hits.Hits {
		head, err1 := uses.ConvertUserToFloatingHead(eclient, element.Id)
		if err1 != nil {
			err = errors.New("there was one or more problems loading results")
			continue
		}
		results = append(results, head)
	}

	return results, err
}
