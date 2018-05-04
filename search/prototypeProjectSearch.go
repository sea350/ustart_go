package search

import (
	"context"
	"errors"
	"fmt"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//PrototypeProjectSearch ... Attempt at fully functional project search, returns Floatinghead
/* Inputs:
eclient		-> ???
sortBy 		-> 0: Relevance, 1: Popularity, 2: Newest/Age
searchBy 	-> [Name, Username] if true will add respective filter
mustMajor 	-> Array of Majors that each result must have
mustTag 	-> Array of Tags that each result must have
mustLoc 	-> Location that the result must have
searchTerm 	-> The term user is inputting
*/
func PrototypeProjectSearch(eclient *elastic.Client, searchTerm string, sortBy int, searchBy []bool, mustMajor []string, mustTag []string, mustLoc []types.LocStruct) ([]types.FloatingHead, error) {
	ctx := context.Background()

	var results []types.FloatingHead
	var searchArr []string
	query := elastic.NewBoolQuery()

	stringArray := strings.Split(searchTerm, `","`)
	for _, element := range stringArray {
		searchArr = append(searchArr, strings.ToLower(element))
	}

	if len(searchBy) >= 4 {
		//Name
		if searchBy[0] {
			query = uses.MultiWildCardQuery(query, "Name", stringArray, true)
		}
		//URLName
		if searchBy[1] {
			query = uses.MultiWildCardQuery(query, "URLName", stringArray, true)
		}
		//Tags
		if searchBy[2] {
			query = uses.MultiWildCardQuery(query, "Tags", stringArray, true)
		}
		//ListNeeded
		if searchBy[3] {
			query = uses.MultiWildCardQuery(query, "ListNeeded", stringArray, true)
		}
	} else {
		fmt.Println("WARNING: searchBy array is too short")
	}
	// Major
	if len(mustMajor) > 0 {
		for _, element := range mustMajor {
			//Check if NewMatchQuery order is correct
			query = query.Must(elastic.NewMatchQuery("ListNeeded", element))
		}
	}
	// Tag
	if len(mustTag) > 0 {
		for _, element := range mustTag {
			//Check if NewMatchQuery order is correct
			query = query.Must(elastic.NewMatchQuery("Tags", element))
		}
	}

	searchResults, err := eclient.Search().
		Index(globals.ProjectIndex).
		Query(query).
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
		head, err1 := uses.ConvertProjectToFloatingHead(eclient, element.Id)
		if err1 != nil {
			err = errors.New("there was one or more problems loading results")
			continue
		}
		results = append(results, head)
	}

	return results, err
}
