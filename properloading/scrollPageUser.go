package properloading

import (
	"context"
	"io"
	"log"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
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
	//set up user query
	usrQuery := elastic.NewBoolQuery()
	usrQuery = usrQuery.Must(elastic.NewTermQuery("PosterID", docID))
	usrQuery = usrQuery.Should(elastic.NewTermQuery("Classification", "0"))
	//usrQuery = usrQuery.Should(elastic.NewTermQuery("Classification", "2"))

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
	if err == io.EOF {
		return "", arrResults, 0, err //we might need special treatment for EOF error
	}
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return "", arrResults, 0, err
	}

	//fmt.Println(res.Hits.TotalHits)

	/*
		for _, hit := range res.Hits.Hits {
			head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, false)
			arrResults = append(arrResults, head)
			if err != nil {
				return res.ScrollId, arrResults, int(res.Hits.TotalHits), errors.New("ISSUE WITH CONVERT FUNCTION")

			}

			if err == io.EOF {
				return res.ScrollId, arrResults, int(res.Hits.TotalHits), errors.New("Out of bounds")

			}

		}
	*/

	return res.ScrollId, arrResults, int(res.Hits.TotalHits), err
}
