package search

import (
	"context"
	"fmt"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//BoolUserSearch ... Search user with a bool query taking in username and filter
func BoolUserSearch(eclient *elastic.Client, input string) ([]types.FloatingHead, error) {

	//split input
	inputList := splitInput(input)

	//initialize bool query
	boolSearch := elastic.NewBoolQuery()

	//adding refinement to Bquery
	for _, item := range inputList {
		boolSearch = boolSearch.Should(
			elastic.NewTermQuery("Username", item),
			elastic.NewTermQuery("FirstName", item),
			elastic.NewTermQuery("LastName", item))
	}

	//execute query
	searchResults, err := eclient.Search().
		Index(globals.UserIndex).
		Query(boolSearch).
		Do(context.Background())

	//storing and converting result to floatinghead type
	var results []types.FloatingHead
	for _, element := range searchResults.Hits.Hits {
		head, err := uses.ConvertUserToFloatingHead(eclient, element.Id, ``)
		if err != nil {
			fmt.Println("err: search/boolUserSearch line 46 index ")
		}
		results = append(results, head)
	}

	return results, err
}

func splitInput(input string) []string {
	//separated by spaces
	return strings.Split(input, " ")
}
