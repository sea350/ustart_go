package search

import (
	"context"
	"fmt"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
	//get "github.com/sea350/ustart_go/get"
	//"encoding/json"
	//"errors"
)

//UserBasic ...
//basic user search
func UserBasic(eclient *elastic.Client, searchTerm string) ([]types.FloatingHead, error) {
	ctx := context.Background()

	newMatchQuery := elastic.NewMatchQuery("Username", searchTerm)

	var results []types.FloatingHead
	searchResults, err := eclient.Search().
		Index(globals.UserIndex). // search in index "twitter"
		Query(newMatchQuery).     // specify the query
		Sort("user", true).       // sort by "user" field, ascending
		Pretty(true).             // pretty print request and response JSON
		Do(ctx)                   // execute

	//nResults := searchResults.TotalHits

	for i, element := range searchResults.Hits.Hits {
		head, err := uses.ConvertUserToFloatingHead(eclient, element.Id)
		if err != nil {
			fmt.Println("err: search/userBasic line 36 index ", i)
		}
		results = append(results, head)
	}

	return results, err
}
