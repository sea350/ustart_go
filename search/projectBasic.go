package search

import (
	"context"
	"fmt"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
	//get "github.com/sea350/ustart_go/get"
	//"encoding/json"
	//"errors"
)

//ProjectBasic ...
//basic project search
func ProjectBasic(eclient *elastic.Client, searchTerm string) ([]types.FloatingHead, func() int64, error) {
	ctx := context.Background()

	var results []types.FloatingHead
	newMatchQuery := elastic.NewMatchQuery("Name", searchTerm)

	searchResults, err := eclient.Search().
		Index(globals.ProjectIndex). // search in index "twitter"
		Query(newMatchQuery).        // specify the query
		//Sort("user", true).       // sort by "user" field, ascending
		Pretty(true). // pretty print request and response JSON
		Do(ctx)       // execute

	nHits := searchResults.TotalHits

	if err != nil {
		return results, nHits, err
	}

	for i, element := range searchResults.Hits.Hits {
		head, err := uses.ConvertProjectToFloatingHead(eclient, element.Id)
		if err != nil {
			fmt.Println("err: search/project line 47 index ", i)
		}
		results = append(results, head)
	}

	return results, nHits, err

}
