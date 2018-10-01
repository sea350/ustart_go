package properloading

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/backend/globals"
	types "github.com/sea350/ustart_go/backend/types"
	"github.com/sea350/ustart_go/backend/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ScrollPageProject ...
//Scrolls through docs being loaded on project page
func ScrollPageProject(eclient *elastic.Client, docID string, viewerID string, scrollID string) (string, []types.JournalEntry, int, error) {

	ctx := context.Background()

	//set up project query
	projQuery := elastic.NewBoolQuery()
	projQuery = projQuery.Must(elastic.NewTermsQuery("ReferenceID", strings.ToLower(docID)))
	projQuery = projQuery.Must(elastic.NewTermsQuery("Classification", 3, 5))
	projQuery = projQuery.Must(elastic.NewTermQuery("Visible", true))

	//yeah....

	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(projQuery).
		Sort("TimeStamp", false).
		Size(5)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)
	if err == io.EOF {
		return "", arrResults, 0, err //we might need special treatment for EOF error
	}
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return "", arrResults, 0, err
	}

	for _, hit := range res.Hits.Hits {
		// fmt.Println(hit.Id)
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, viewerID, true)
		arrResults = append(arrResults, head)
		if err != nil {
			return res.ScrollId, arrResults, int(res.Hits.TotalHits), errors.New("ISSUE WITH CONVERT FUNCTION")
		}
	}

	return res.ScrollId, arrResults, int(res.Hits.TotalHits), err
}
