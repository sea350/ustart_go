package search

import (
	"context"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
	//get "github.com/sea350/ustart_go/get"
	//"encoding/json"
	//"errors"
)

type ProjectResult struct {
	Name    string `json:"Name"`
	Avatar  string `json:"Avatar"`
	URLName string `json:"URLName"`
}

//Project ...
//basic project search
func Project(eclient *elastic.Client, field string, searchTerm string) ([]ProjectResult, func() int64, error) {
	ctx := context.Background()

	var results []ProjectResult
	newMatchQuery := elastic.NewMatchQuery(field, searchTerm)

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

	for _, element := range searchResults.Hits.Hits {
		var newResult ProjectResult
		proj, err := get.ProjectByID(eclient, element.Id)

		newResult.Name = proj.Name
		newResult.Avatar = proj.Avatar
		newResult.URLName = proj.URLName

		results = append(results, newResult)
	}

	return results, nHits, err

}
