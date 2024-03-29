package properloading

import (
	"context"
	"errors"
	"io"
	"log"

	elastic "github.com/olivere/elastic"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//ScrollPageDash ...
//Scrolls through docs being loaded
func ScrollPageDash(eclient *elastic.Client, docIDs []string, viewerID string, scrollID string) (string, []types.JournalEntry, int, error) {

	ctx := context.Background()

	ids := make([]interface{}, 0)
	for id := range docIDs {

		// var dash = rune('-')
		// var underscore = rune('_')
		// var tempRuneArr []rune
		// for _, char := range docIDs[id] {
		// 	if char != dash && char != underscore {
		// 		tempRuneArr = append(tempRuneArr, char)
		// 	}
		// }
		// docIDs[id] = string(tempRuneArr)

		ids = append([]interface{}{docIDs[id]}, ids...)
	}

	//set up user query
	usrQuery := elastic.NewBoolQuery()
	usrQuery = usrQuery.Must(elastic.NewTermsQuery("PosterID.keyword", ids...))
	usrQuery = usrQuery.Must(elastic.NewTermsQuery("Classification", 0, 2))
	usrQuery = usrQuery.Must(elastic.NewTermQuery("Visible", true))

	//set up project query
	projQuery := elastic.NewBoolQuery()
	projQuery = projQuery.Must(elastic.NewTermsQuery("ReferenceID.keyword", ids...))
	projQuery = projQuery.Must(elastic.NewTermsQuery("Classification", 3, 5))
	projQuery = projQuery.Must(elastic.NewTermQuery("Visible", true))
	//yeah....
	finalQuery := elastic.NewBoolQuery()
	finalQuery = finalQuery.Should(usrQuery)
	finalQuery = finalQuery.Should(projQuery)
	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(finalQuery).
		Sort("TimeStamp", false).
		Size(20)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)

	if (err != nil && err != io.EOF) || res == nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return "", arrResults, 0, err
	}

	var report error
	for _, hit := range res.Hits.Hits {
		// fmt.Println(hit.Id)
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, viewerID, true)
		if err != nil {
			report = errors.New("One or more problems loading journal entries")
			continue
		}
		arrResults = append(arrResults, head)

	}

	return res.ScrollId, arrResults, int(res.Hits.TotalHits), report
}
