package search

import (
	"context"
	"encoding/json"
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

//PrototypeProjectSearchScroll ... Attempt at fully functional project search, returns Floatinghead
/* Inputs:
eclient		-> ???
sortBy 		-> 0: Relevance, 1: Popularity, 2: Newest/Age
searchBy 	-> [Name, Username] if true will add respective filter
mustMajor 	-> Array of Majors that each result must have
mustTag 	-> Array of Tags that each result must have
mustLoc 	-> Location that the result must have
searchTerm 	-> The term user is inputting
*/
func PrototypeProjectSearchScroll(eclient *elastic.Client, searchTerm string, sortBy int, searchBy []bool, mustMajor []string, mustTag []string, mustLoc []types.LocStruct, scrollID string) (int, string, []types.FloatingHead, error) {
	ctx := context.Background()

	var results []types.FloatingHead
	var stringArray []string
	query := elastic.NewBoolQuery()

	err := json.Unmarshal([]byte(searchTerm), &stringArray)
	if err != nil {
		return 0, "", results, err
	}

	query = query.Must(elastic.NewTermQuery("Visible", true))

	if len(searchBy) >= 4 {
		//Name
		if searchBy[0] {
			query = uses.MultiWildCardQuery(query, "Name", stringArray, true)
			for _, element := range stringArray {
				query = query.Should(elastic.NewFuzzyQuery("Name", strings.ToLower(element)).Fuzziness(2))
			}
		}
		//URLName
		if searchBy[1] {
			query = uses.MultiWildCardQuery(query, "URLName", stringArray, true)
			for _, element := range stringArray {
				query = query.Should(elastic.NewFuzzyQuery("URLName", strings.ToLower(element)).Fuzziness(2))
			}
		}
		//Tags
		if searchBy[2] {
			query = uses.MultiWildCardQuery(query, "Tags", stringArray, true)
			for _, element := range stringArray {
				query = query.Should(elastic.NewFuzzyQuery("Tags", strings.ToLower(element)).Fuzziness(1))
			}
		}
		//ListNeeded
		if searchBy[3] {
			query = uses.MultiWildCardQuery(query, "ListNeeded", stringArray, true)
			for _, element := range stringArray {
				query = query.Should(elastic.NewFuzzyQuery("ListNeeded", strings.ToLower(element)).Fuzziness(1))
			}
		}
	} else {
		fmt.Println("WARNING: searchBy array is too short")
	}
	// Major
	if len(mustMajor) > 0 {
		for _, element := range mustMajor {
			//Check if NewMatchQuery order is correct
			query = query.Must(elastic.NewMatchQuery("ListNeeded", element))
			for _, element := range stringArray {
				query = query.Should(elastic.NewFuzzyQuery("ListNeeded", strings.ToLower(element)).Fuzziness(1))
			}
		}
	}
	// Tag
	if len(mustTag) > 0 {
		tags := make([]interface{}, 0)
		for tag := range mustTag {
			tags = append([]interface{}{strings.ToLower(mustTag[tag])}, tags...)
		}

		query = query.Must(elastic.NewTermsQuery("Tags", tags...))
		// for _, element := range mustTag {

		// 	//Check if NewMatchQuery order is correct
		// 	query = query.Must(elastic.NewMatchQuery("Tags", element))
		// 	for _, element := range stringArray {
		// 		query = query.Should(elastic.NewFuzzyQuery("Tags", strings.ToLower(element)).Fuzziness(1))
		// 	}
		// }
	}

	scroll := eclient.Scroll().
		Index(globals.ProjectIndex).
		Query(query).
		Size(10)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)
	if !(err == io.EOF && res != nil) && err != nil {
		if err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		return 0, "", results, err
	}

	for _, element := range res.Hits.Hits {
		head, err1 := uses.ConvertProjectToFloatingHead(eclient, element.Id)
		if err1 != nil {
			err = errors.New("there was one or more problems loading results")
			continue
		}
		results = append(results, head)
	}

	return int(res.Hits.TotalHits), res.ScrollId, results, err
}
