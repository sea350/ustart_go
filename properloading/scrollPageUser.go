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
