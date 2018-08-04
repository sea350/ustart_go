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

//ScrollPageUser ...
//Scrolls through docs being loaded on the user wall
func ScrollPageUser(eclient *elastic.Client, docIDs []string, scrollID string) (string, []types.JournalEntry, int64, error) {

	ctx := context.Background()

	ids := make([]interface{}, 0)
	for id := range docIDs {
		ids = append([]interface{}{strings.ToLower(docIDs[id])}, ids...)
	}

	//set up user query
	usrQuery := elastic.NewBoolQuery()
	usrQuery = usrQuery.Must(elastic.NewTermsQuery("PosterID", ids...))
	usrQuery = usrQuery.Should(elastic.NewTermQuery("Classification", "0"))
	usrQuery = usrQuery.Should(elastic.NewTermQuery("Classification", "2"))

	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(usrQuery).
		Sort("TimeStamp", false).
		Size(5)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)

	for _, hit := range res.Hits.Hits {
		// fmt.Println(hit.Id)
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
