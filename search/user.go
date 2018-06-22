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

//User ...
//basic user search
func User(eclient *elastic.Client, field string, searchTerm string) ([]types.User, error) {
	ctx := context.Background()

	newMatchQuery := elastic.NewMatchQuery(field, searchTerm)

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
			fmt.Println("err: search/user line 43 index ", i)
		}
		results = append(results, head)
	}

	var ret []types.User
	return ret, err

}
