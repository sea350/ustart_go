package search

import (
	"context"
	"fmt"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

//SearchProfile ... Attempt at fully functional profile search, returns Floatinghead
func SearchProfile(eclient *elastic.Client, searchTerm string) ([]types.FloatingHead, error) {
	ctx := context.Background()
	var results []types.FloatingHead

	newMatchQuery := elastic.NewMultiMatchQuery(searchTerm, "FirstName", "LastName", "UserName", "Major", "Tags")
	searchResults, err := eclient.Search().
		Index(globals.UserIndex).
		Query(newMatchQuery).
		Pretty(true).
		Do(ctx)

	if err != nil {
		fmt.Println("waduhek")
	}

	//Testing outputs
	numHits := searchResults.Hits.TotalHits
	fmt.Println("Number of Hits: ", numHits)
	for _, s := range searchResults.Hits.Hits {
		u, _ := get.UserByID(eclient, s.Id)
		fmt.Println(u.FirstName, u.LastName)
	}

	for i, element := range searchResults.Hits.Hits {
		head, err := uses.ConvertUserToFloatingHead(eclient, element.Id)
		if err != nil {
			fmt.Println("error search/userBasic line 43 index", i)
		}
		results = append(results, head)
	}

	return results, err
}
