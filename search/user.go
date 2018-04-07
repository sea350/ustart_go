package search

import (
	"context"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
	//get "github.com/sea350/ustart_go/get"
	//"encoding/json"
	//"errors"
)

type UserResult struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Avatar    string `json:"Avatar"`
	Username  string `json:"Username"`
}

//User ...
//basic user search
func User(eclient *elastic.Client, field string, searchTerm string) ([]types.User, error) {
	ctx := context.Background()

	newMatchQuery := elastic.NewMatchQuery(field, searchTerm)

	var results []UserResult
	searchResults, err := eclient.Search().
		Index(globals.UserIndex). // search in index "twitter"
		Query(newMatchQuery).     // specify the query
		Sort("user", true).       // sort by "user" field, ascending
		Pretty(true).             // pretty print request and response JSON
		Do(ctx)                   // execute

	nResults := searchResults.TotalHits

	for _, element := range searchResults.Hits.Hits {
		var newResult UserResult
		usr, err := get.UserByID(eclient, element.Id)

		newResult.FirstName = usr.FirstName
		newResult.LastName = usr.LastName
		newResult.Avatar = usr.Avatar
		newResult.Username = usr.Username

		results = append(results, newResult)
	}

	var ret []types.User
	return ret, err

}
