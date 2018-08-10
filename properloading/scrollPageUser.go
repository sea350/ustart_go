package properloading

import (
	"context"
	"errors"
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

	//set up user query
	usrQuery := elastic.NewBoolQuery()
	usrQuery = usrQuery.Must(elastic.NewTermQuery("PosterID", strings.ToLower(docID)))
	usrQuery = usrQuery.Should(elastic.NewTermQuery("Classification", 0))

	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(usrQuery).
		Sort("TimeStamp", false).
		Size(12)

	scroll = scroll.ScrollId(scrollID)

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("scrollID: " + scrollID)

	res, err := scroll.Do(ctx)
	if err != nil {
		return "", arrResults, 0, err
	}

	//fmt.Println(res.Hits.TotalHits)

	for _, hit := range res.Hits.Hits {
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, true)
		arrResults = append(arrResults, head)
		if err != nil {
			return res.ScrollId, arrResults, int(res.Hits.TotalHits), errors.New("ISSUE WITH CONVERT FUNCTION")

		}
	}

	return res.ScrollId, arrResults, int(res.Hits.TotalHits), err
}
