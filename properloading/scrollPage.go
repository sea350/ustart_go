package properloading

import (
	"context"
	"fmt"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ScrollPage ...
//Scrolls through docs being loaded
func ScrollPage(eclient *elastic.Client, docIDs []string, scrollID string) (string, []types.JournalEntry, error) {

	ctx := context.Background()

	var searchThese []string
	for id := range docIDs {
		searchThese = append(searchThese, strings.ToLower(docIDs[id]))
	}
	query := elastic.NewTermsQuery("PosterID", searchThese[0], searchThese[1])
	fmt.Println(searchThese)
	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(query).
		Size(2)

	/*
		if scrollID != "" {
			scroll = scroll.ScrollId(scrollID)
		}*/

	res, err := scroll.Do(ctx)
	fmt.Println("RES HITS:", res.TotalHits())
	if err != nil {
		return "res.ScrollID", arrResults, err
	}
	// fmt.Println(res)
	// for _, hit := range res.Hits.Hits {
	// 	fmt.Println(hit.Id)
	// 	head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, false)
	// 	arrResults = append(arrResults, head)
	// 	if err != nil {
	// 		return res.ScrollId, arrResults, errors.New("ISSUE WITH CONVERT FUNCTION")

	// 	}

	// 	if err == io.EOF {
	// 		return res.ScrollId, arrResults, errors.New("Out of bounds")

	// 	}

	// }
	return res.ScrollId, arrResults, err
}
