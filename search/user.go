package search

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
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

	searchResults, err := eclient.Search().
		Index(globals.UserIndex). // search in index "twitter"
		Query(newMatchQuery).     // specify the query
		Sort("user", true).       // sort by "user" field, ascending
		Pretty(true).             // pretty print request and response JSON
		Do(ctx)                   // execute

	nResults := searchResults.TotalHits

	var ret []types.User
	return ret, err

}
