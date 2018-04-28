package search

import (
	"context"
	"errors"
	"fmt"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//PrototypeUserSearch ... Attempt at fully functional user search, returns Floatinghead
/* Inputs:
eclient		-> ???
sortBy 		-> 0: Relevance, 1: Popularity, 2: Newest/Age
searchBy 	-> [Name, Username] if true will add respective filter
mustMajor 	-> Array of Majors that each result must have
mustTag 	-> Array of Tags that each result must have
mustLoc 	-> Location that the result must have
searchTerm 	-> The term user is inputting
*/
func PrototypeUserSearch(eclient *elastic.Client, searchTerm string, sortBy int, searchBy []bool, mustMajor []string, mustTag []string, mustLoc []types.LocStruct) ([]types.FloatingHead, error) {
	ctx := context.Background()

	var results []types.FloatingHead

	//, "Description", "URLName", "Tags"
	// newMatchQuery := elastic.NewMultiMatchQuery(searchTerm, "FirstName", "LastName")
	newMatchQuery := elastic.NewBoolQuery()

	if len(searchBy) >= 3 {
		//Name
		if searchBy[0] {
			newMatchQuery = newMatchQuery.Should(elastic.NewWildcardQuery("FirstName", `*`+searchTerm+`*`))
			newMatchQuery = newMatchQuery.Should(elastic.NewWildcardQuery("LastName", `*`+searchTerm+`*`))
		}
		//Username
		if searchBy[1] {
			newMatchQuery = newMatchQuery.Should(elastic.NewWildcardQuery("Username", `*`+searchTerm+`*`))
		}
		//Tags
		if searchBy[2] {
			newMatchQuery = newMatchQuery.Should(elastic.NewWildcardQuery("Tags", `*`+searchTerm+`*`))
		}
	} else {
		fmt.Println("WANRING: searchBy array is too short")
	}
	// Major
	if len(mustMajor) > 0 {
		for _, element := range mustMajor {
			//Check if NewMatchQuery order is correct
			newMatchQuery = newMatchQuery.Must(elastic.NewMatchQuery("Majors", element))
		}
	}
	// Tag
	if len(mustTag) > 0 {
		for _, element := range mustTag {
			//Check if NewMatchQuery order is correct
			newMatchQuery = newMatchQuery.Must(elastic.NewMatchQuery("Tags", element))
		}
	}

	searchResults, err := eclient.Search().
		Index(globals.UserIndex).
		Query(newMatchQuery).
		Pretty(true).
		Do(ctx)

	if err != nil {
		return results, err
	}

	// Testing Outputs
	// fmt.Println("Number of Hits: ", searchResults.Hits.TotalHits)
	// for _, s := range searchResults.Hits.Hits {
	// 	u, _ := get.UserByID(eclient, s.Id)
	// 	// fmt.Println(u.FirstName, u.LastName)
	// 	fmt.Println(u.FirstName, u.LastName)
	// }

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
