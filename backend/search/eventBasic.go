package search

import (
	"context"
	"fmt"

	globals "github.com/sea350/ustart_go/backend/globals"
	"github.com/sea350/ustart_go/backend/types"
	"github.com/sea350/ustart_go/backend/uses"
	elastic "gopkg.in/olivere/elastic.v5"
	//get "github.com/sea350/ustart_go/backend/get"
	//"encoding/json"
	//"errors"
)

//EventBasic ...
//basic event search
func EventBasic(eclient *elastic.Client, searchTerm string) ([]types.FloatingHead, func() int64, error) {
	ctx := context.Background()

	var results []types.FloatingHead
	newMatchQuery := elastic.NewMatchQuery("Name", searchTerm)

	searchResults, err := eclient.Search().
		Index(globals.EventIndex). // search in index "twitter"
		Query(newMatchQuery).      // specify the query
		//Sort("user", true).       // sort by "user" field, ascending
		Pretty(true). // pretty print request and response JSON
		Do(ctx)       // execute

	nHits := searchResults.TotalHits

	if err != nil {
		return results, nHits, err
	}

	for i, element := range searchResults.Hits.Hits {
		head, err := uses.ConvertEventToFloatingHead(eclient, element.Id)
		if err != nil {
			fmt.Println("err: search/event line 47 index ", i)
		}
		results = append(results, head)
	}

	return results, nHits, err

}
