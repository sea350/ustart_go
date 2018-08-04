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

//ScrollPageEvents ...
//Scrolls through event docs being loaded
func ScrollPageEvents(eclient *elastic.Client, docIDs []string, scrollID string) (string, []types.JournalEntry, int64, error) {

	ctx := context.Background()

	docIDModified := make([]interface{}, 0)
	for id := range docIDs {
		docIDModified = append([]interface{}{strings.ToLower(docIDs[id])}, docIDModified...)
	}

	//set up event query
	evntQuery := elastic.NewBoolQuery()
	evntQuery = evntQuery.Must(elastic.NewTermsQuery("ReferenceID", docIDModified...))
	evntQuery = evntQuery.Should(elastic.NewTermQuery("Classification", "6"))
	//evntQuery = evntQuery.Should(elastic.NewTermQuery("Classification", "7"))
	//evntQuery = evntQuery.Should(elastic.NewTermQuery("Classification", "8"))

	var arrResults []types.JournalEntry
	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(evntQuery).
		Sort("TimeStamp", false).
		Size(5)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)

	for _, hit := range res.Hits.Hits {
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, false)
		arrResults = append(arrResults, head)
		if err != nil {
			return res.ScrollId, arrResults, res.Hits.TotalHits, errors.New("ISSUE WITH CONVERT FUNCTION")

		}

		if err == io.EOF {
			return res.ScrollId, arrResults, res.Hits.TotalHits, errors.New("Out of bounds")

		}
	}

	return res.ScrollId, arrResults, res.Hits.TotalHits, err
}
