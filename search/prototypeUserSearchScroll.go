package search

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//PrototypeUserSearchScroll ... Attempt at fully functional user search, returns Floatinghead
/* Inputs:
eclient		-> ???
sortBy 		-> 0: Relevance, 1: Popularity, 2: Newest/Age
searchBy 	-> [Name, Username] if true will add respective filter
mustMajor 	-> Array of Majors that each result must have
mustTag 	-> Array of Tags that each result must have
mustLoc 	-> Location that the result must have
searchTerm 	-> The term user is inputting
*/
func PrototypeUserSearchScroll(eclient *elastic.Client, searchTerm string, sortBy int, searchBy []bool, mustMajor []string, mustTag []string, mustLoc []types.LocStruct, scrollID string) (int, string, []types.FloatingHead, error) {
	ctx := context.Background()
	var results []types.FloatingHead
	var searchArr []string
	query := elastic.NewBoolQuery()
	searchArr = strings.Split(searchTerm, ` `)

	query = query.Must(elastic.NewTermQuery("Verified", true))

	if len(searchBy) >= 3 {
		//Name
		if searchBy[0] {
			query = uses.MultiWildCardQuery(query, "FirstName", searchArr, true)
			query = uses.MultiWildCardQuery(query, "LastName", searchArr, true)

			for _, element := range searchArr {
				query = query.Should(elastic.NewFuzzyQuery("FirstName", strings.ToLower(element)).Fuzziness(1))
				query = query.Should(elastic.NewFuzzyQuery("LastName", strings.ToLower(element)).Fuzziness(1))
			}
		}
		//Username
		if searchBy[1] {
			query = uses.MultiWildCardQuery(query, "Username", searchArr, true)

			for _, element := range searchArr {
				query = query.Should(elastic.NewFuzzyQuery("Username", strings.ToLower(element)).Fuzziness(1))
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("WARNING: searchBy array is too short")
	}
	// Major
	if len(mustMajor) > 0 {
		for _, element := range mustMajor {
			//Check if NewMatchQuery order is correct
			query = query.Must(elastic.NewMatchQuery("Majors", strings.ToLower(element)))
			query = query.Should(elastic.NewFuzzyQuery("Majors", strings.ToLower(element)).Fuzziness(1))
		}
	}

	if len(mustTag) > 0 {
		for _, tag := range mustTag {
			query = query.Must(elastic.NewTermQuery("Tags", strings.ToLower(tag)))
		}
	}

	scroll := eclient.Scroll().
		Index(globals.UserIndex).
		Query(query).
		Size(10)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)
	if err == io.EOF {
		return 0, "", results, err
	}
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	for _, element := range res.Hits.Hits {
		head, err1 := uses.ConvertUserToFloatingHead(eclient, element.Id)
		if err1 != nil {
			err = errors.New("there was one or more problems loading results")
			continue
		}
		results = append(results, head)
	}

	fmt.Println(int(res.Hits.TotalHits), res.ScrollId, results, err)

	return int(res.Hits.TotalHits), res.ScrollId, results, err
}
