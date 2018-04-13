package search

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//SearchProject ... Attempt at fully functional project search, returns Floatinghead
func SearchProject(eclient *elastic.Client, searchTerm string) ([]types.FloatingHead, error) {
	ctx := context.Background()
	var results []types.FloatingHead

	newMatchQuery := elastic.NewMultiMatchQuery(searchTerm, "Name", "Description", "URLName", "Tags")
	searchResults, err := eclient.Search().
		Index(globals.ProjectIndex).
		Query(newMatchQuery).
		Pretty(true).
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
