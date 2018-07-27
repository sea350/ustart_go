package properloading

import (
	"context"
	"errors"
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

	tmp := make([]interface{}, 0)
	for id := range docIDs {
		tmp = append([]interface{}{strings.ToLower(docIDs[id])}, tmp...)
	}

	// for id := range docIDs {
	// 	searchThese[id] = strings.ToLower(docIDs[id])
	// }
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermsQuery("PosterID", tmp...))
	// query = query.Must(elastic.NewRangeQuery("TimeStamp").From("2017-01-01").To("2018-04-19").Boost(3))
	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(query).
		Sort("TimeStamp", true).
		Size(10)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)

	for _, hit := range res.Hits.Hits {
		// fmt.Println(hit.Id)
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, false)
		arrResults = append(arrResults, head)
		if err != nil {
			return res.ScrollId, arrResults, errors.New("ISSUE WITH CONVERT FUNCTION")

		}

		if err == io.EOF {
			return res.ScrollId, arrResults, errors.New("Out of bounds")

		}

	}

	return res.ScrollId, arrResults, err
}
