package properloading

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ScrollPageDash ...
//Scrolls through docs being loaded
func ScrollPageDash(eclient *elastic.Client, docIDs []string, viewerID string, scrollID string) (string, []types.JournalEntry, int, error) {

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

	//set up project query
	projQuery := elastic.NewBoolQuery()
	projQuery = projQuery.Should(elastic.NewTermQuery("Classification", "3"))
	projQuery = projQuery.Must(elastic.NewTermsQuery("ReferenceID", ids...))
	//yeah....
	finalQuery := usrQuery.Should(projQuery)
	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(finalQuery).
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
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, viewerID, false)
		arrResults = append(arrResults, head)
		if err != nil && err != errors.New("This entry is not visible") {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("ISSUE WITH CONVERT FUNCTION")
			return res.ScrollId, arrResults, int(res.Hits.TotalHits), err
		}

		if err == io.EOF {
			return res.ScrollId, arrResults, int(res.Hits.TotalHits), errors.New("Out of bounds")

		}

	}

	return res.ScrollId, arrResults, int(res.Hits.TotalHits), err
}
