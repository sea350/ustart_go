package search

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//PrototypeUserScrollSearch ... Attempt at fully functional user search, returns Floatinghead
/* Inputs:
eclient		-> ???
sortBy 		-> 0: Relevance, 1: Popularity, 2: Newest/Age
searchBy 	-> [Name, Username] if true will add respective filter
mustMajor 	-> Array of Majors that each result must have
mustTag 	-> Array of Tags that each result must have
mustLoc 	-> Location that the result must have
searchTerm 	-> The term user is inputting
*/
func PrototypeUserSearchScroll(eclient *elastic.Client, searchTerm string, sortBy int, searchBy []bool, mustMajor []string, mustTag []string, mustLoc []types.LocStruct, scrollID string) (string, []types.FloatingHead, error) {
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

	/*
		searchResults, err := eclient.Search().
			Index(globals.UserIndex).
			Query(query).
			Pretty(true).
			Do(ctx)

		if err != nil {
			return results, err
		}
	*/

	scroll := eclient.Scroll().
		Index(globals.UserIndex).
		Query(query).
		Size(2)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}
	// Testing Outputs
	// fmt.Println("Number of Hits: ", searchResults.Hits.TotalHits)
	// for _, s := range searchResults.Hits.Hits {
	// 	u, _ := get.UserByID(eclient, s.Id)
	// 	// fmt.Println(u.FirstName, u.LastName)
	// 	fmt.Println(u.FirstName, u.LastName)
	// }

	res, err := scroll.Do(ctx)
	if err == io.EOF {
		return res.ScrollId, results, err
	}
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	fmt.Println("\n\n", res.Hits.TotalHits, "\n\n")
	if res.Hits.TotalHits > 0 {
		for _, element := range res.Hits.Hits {
			head, err1 := uses.ConvertUserToFloatingHead(eclient, element.Id)
			if err1 != nil {
				err = errors.New("there was one or more problems loading results")
				continue
			}
			results = append(results, head)
		}
	}

	return res.ScrollId, results, err
}
