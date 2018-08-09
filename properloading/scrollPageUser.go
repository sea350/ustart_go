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

//ScrollPageUser ...
//Scrolls through docs being loaded on the user wall
func ScrollPageUser(eclient *elastic.Client, docID string, scrollID string) (string, []types.JournalEntry, int, error) {

	ctx := context.Background()

	/*
		ids := make([]interface{}, 0)
		for id := range docIDs {
			ids = append([]interface{}{strings.ToLower(docIDs[id])}, ids...)
		}
	*/

	//Mimics the above, but for a single entry
	userID := strings.ToLower(docID)

	//set up user query
	usrQuery := elastic.NewBoolQuery()
	usrQuery = usrQuery.Must(elastic.NewTermQuery("PosterID", userID))
	usrQuery = usrQuery.Should(elastic.NewTermQuery("Classification", "0"))
	usrQuery = usrQuery.Should(elastic.NewTermQuery("Classification", "2"))

	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(usrQuery).
		Sort("TimeStamp", false).
		Size(12)

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

	//fmt.Println(res.Hits.TotalHits)

	for _, hit := range res.Hits.Hits {
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, true)
		arrResults = append(arrResults, head)
		if err != nil {
			return res.ScrollId, arrResults, int(res.Hits.TotalHits), errors.New("ISSUE WITH CONVERT FUNCTION")

		}

		if err == io.EOF {
			return res.ScrollId, arrResults, int(res.Hits.TotalHits), errors.New("Out of bounds")

		}

	}

	return res.ScrollId, arrResults, int(res.Hits.TotalHits), err
}
