package search

import (
	"context"
	"errors"
	"fmt"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//SearchProject ... Attempt at fully functional project search, returns Floatinghead
func SearchProject(eclient *elastic.Client, searchTerm string) ([]types.FloatingHead, error) {
	fmt.Println("STARTING")
	ctx := context.Background()
	var results []types.FloatingHead

	//, "Description", "URLName", "Tags"
	newMatchQuery := elastic.NewMultiMatchQuery(searchTerm, "Name", "Tags", "URLName")
	searchResults, err := eclient.Search().
		Index(globals.ProjectIndex).
		Query(newMatchQuery).
		Pretty(true).
		Do(ctx)

	if err != nil {
		return results, err
	}
	// //Testing outputs
	// fmt.Println("ERROR:", err, "\n")
	// numHits := searchResults.Hits.TotalHits
	// fmt.Println("Number of Hits: ", numHits)
	// for _, s := range searchResults.Hits.Hits {
	// 	u, _ := get.ProjectByID(eclient, s.Id)
	// 	fmt.Println(u.Name, u.URLName)
	// }
	// return results, err

	// if err != nil {
	// 	fmt.Println("Waduhek\n")
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
