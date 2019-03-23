package properloading

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	elastic "github.com/olivere/elastic"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//ScrollPageProject ...
//Scrolls through docs being loaded on project page
func ScrollPageProject(eclient *elastic.Client, docID string, viewerID string, scrollID string) (string, []types.JournalEntry, int, error) {

	ctx := context.Background()

	// var dash = rune('-')
	// var underscore = rune('_')
	// var tempRuneArr []rune
	// for _, char := range docID {
	// 	if char != dash && char != underscore {
	// 		tempRuneArr = append(tempRuneArr, char)
	// 	}
	// }
	// trimmedID := string(tempRuneArr)

	//set up project query
	projQuery := elastic.NewBoolQuery()
	projQuery = projQuery.Must(elastic.NewTermQuery("ReferenceID.keyword", docID))
	projQuery = projQuery.Must(elastic.NewTermsQuery("Classification", 3, 5))
	projQuery = projQuery.Must(elastic.NewTermQuery("Visible", true))

	//yeah....

	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(projQuery).
		Sort("TimeStamp", false).
		Size(12)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)
	if !(err == io.EOF && res != nil) && err != nil {
		if err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		return "", arrResults, 0, err
	}

	var tempEntry types.Entry
	for _, hit := range res.Hits.Hits {
		err := json.Unmarshal(*hit.Source, &tempEntry) //unmarshal type RawMessage into user struct
		if err != nil {
			return "", arrResults, 0, err
		}
		if tempEntry.ReferenceID != docID {
			continue
		}
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, viewerID, true)
		arrResults = append(arrResults, head)
		if err != nil {
			return res.ScrollId, arrResults, int(res.Hits.TotalHits), errors.New("ISSUE WITH CONVERT FUNCTION")
		}
	}

	return res.ScrollId, arrResults, int(res.Hits.TotalHits), err
}
