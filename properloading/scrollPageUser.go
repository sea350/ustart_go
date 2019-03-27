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

//ScrollPageUser ...
//Scrolls through docs being loaded on the user wall
func ScrollPageUser(eclient *elastic.Client, docID string, viewerID string, scrollID string) (string, []types.JournalEntry, int, error) {

	// var dash = rune('-')
	// var underscore = rune('_')
	// var tempRuneArr []rune
	// for _, char := range docID {
	// 	if char != dash && char != underscore {
	// 		tempRuneArr = append(tempRuneArr, char)
	// 	}
	// }
	// docID = string(tempRuneArr)

	// tempRuneArr = []rune{}
	// for _, char := range viewerID {
	// 	if char != dash && char != underscore {
	// 		tempRuneArr = append(tempRuneArr, char)
	// 	}
	// }
	// viewerID = string(tempRuneArr)

	ctx := context.Background()

	//set up user query
	usrQuery := elastic.NewBoolQuery()
	usrQuery = usrQuery.Must(elastic.NewTermQuery("PosterID.keyword", docID))
	usrQuery = usrQuery.Must(elastic.NewTermsQuery("Classification", 0, 2))
	usrQuery = usrQuery.Must(elastic.NewTermQuery("Visible", true))

	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(usrQuery).
		Sort("TimeStamp", false).
		Size(12)

	scroll = scroll.ScrollId(scrollID)

	res, err := scroll.Do(ctx)
	if (err != nil && err != io.EOF) || res == nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return "", arrResults, 0, err
	}

	//fmt.Println(res.Hits.TotalHits)

	var report error
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
		if err != nil && err != errors.New("This entry is not visible") {
			report = errors.New("One or more problems loading journal entries")

		}
	}

	return res.ScrollId, arrResults, int(res.Hits.TotalHits), report
}
