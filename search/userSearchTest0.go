package search

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "github.com/olivere/elastic"
)

//ProfileSearch ... Attempt at fully functional profile search, returns Floatinghead
func ProfileSearch(eclient *elastic.Client, searchTerm string) ([]types.FloatingHead, error) {
	ctx := context.Background()
	var results []types.FloatingHead

	newMatchQuery := elastic.NewMultiMatchQuery(searchTerm, "FirstName", "LastName", "UserName", "Major", "Tags")
	searchResults, err := eclient.Search().
		Index(globals.UserIndex).
		Query(newMatchQuery).
		Pretty(true).
		Do(ctx)

	if err != nil {
		return results, err
	}

	//Testing outputs
	// numHits := searchResults.Hits.TotalHits
	// fmt.Println("Number of Hits: ", numHits)
	// for _, s := range searchResults.Hits.Hits {
	// 	u, _ := get.UserByID(eclient, s.Id)
	// 	// fmt.Println(u.FirstName, u.LastName)
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
