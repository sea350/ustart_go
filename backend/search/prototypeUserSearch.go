package search

import (
	"context"
	"errors"
	"fmt"
	"strings"

	globals "github.com/sea350/ustart_go/backend/globals"
	"github.com/sea350/ustart_go/backend/types"
	"github.com/sea350/ustart_go/backend/uses"
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
	var searchArr []string
	query := elastic.NewBoolQuery()

	searchArr = strings.Split(searchTerm, ` `)
	/*
		for _, element := range stringArray {
			searchArr = append(searchArr, strings.ToLower(element))
		}
	*/
	//, "Description", "URLName", "Tags"
	// query := elastic.NewMultiMatchQuery(searchTerm, "FirstName", "LastName")

	if len(searchBy) >= 3 {
		//Name
		if searchBy[0] {
			query = uses.MultiWildCardQuery(query, "FirstName", searchArr, true)
			query = uses.MultiWildCardQuery(query, "LastName", searchArr, true)

			for _, element := range searchArr {
				query = query.Should(elastic.NewFuzzyQuery("FirstName", strings.ToLower(element)).Fuzziness(2))
				query = query.Should(elastic.NewFuzzyQuery("LastName", strings.ToLower(element)).Fuzziness(2))
			}
		}
		//Username
		if searchBy[1] {
			query = uses.MultiWildCardQuery(query, "Username", searchArr, true)

			for _, element := range searchArr {
				query = query.Should(elastic.NewFuzzyQuery("Username", strings.ToLower(element)).Fuzziness(2))
			}
		}
		//Tags
		if searchBy[2] {
			query = uses.MultiWildCardQuery(query, "Tags", searchArr, true)

			for _, element := range searchArr {
				query = query.Should(elastic.NewFuzzyQuery("Tags", strings.ToLower(element)).Fuzziness(1))
			}
		}
	} else {
		fmt.Println("WARNING: searchBy array is too short")
	}
	// Major
	if len(mustMajor) > 0 {
		for _, element := range mustMajor {
			//Check if NewMatchQuery order is correct
			query = query.Must(elastic.NewMatchQuery("Majors", strings.ToLower(element)))
			query = query.Should(elastic.NewFuzzyQuery("Majors", strings.ToLower(element)).Fuzziness(2))
		}
	}
	// Tag
	if len(mustTag) > 0 {
		for _, element := range mustTag {
			//Check if NewMatchQuery order is correct
			query = query.Should(elastic.NewMatchQuery("Tags", strings.ToLower(element)))
			query = query.Should(elastic.NewFuzzyQuery("Tags", strings.ToLower(element)).Fuzziness(1))
		}
	}

	searchResults, err := eclient.Search().
		Index(globals.UserIndex).
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
		head, err1 := uses.ConvertUserToFloatingHead(eclient, element.Id)
		if err1 != nil {
			err = errors.New("there was one or more problems loading results")
			continue
		}
		results = append(results, head)
	}

	return results, err
}
