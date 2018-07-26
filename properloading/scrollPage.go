package properloading

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
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
	query := elastic.NewTermsQuery("PosterID", searchThese)

	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(query).
		Size(2)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)

	for _, hit := range res.Hits.Hits {
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, false)
		arrResults = append(arrResults, head)
		if err != nil {
			fmt.Println("what is going on???")
		}
		fmt.Println(head.FirstName)
	}

	if err == io.EOF {
		return res.ScrollId, arrResults, errors.New("Out of bounds")

	}
	if err != nil {
		fmt.Println(err)
	}

	return res.ScrollId, arrResults, err

}
