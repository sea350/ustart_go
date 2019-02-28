package properloading

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"

	elastic "github.com/olivere/elastic"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//ScrollPageUser ...
//Scrolls through docs being loaded on the user wall
func ScrollPageUser(eclient *elastic.Client, docID string, viewerID string, scrollID string) (string, []types.JournalEntry, int, error) {

	var dash byte = '-'
	var underscore byte = '_'
	for docID[0] == dash || docID[0] == underscore {
		docID = docID[1:]
	}
	for viewerID[0] == dash || viewerID[0] == underscore {
		viewerID = viewerID[1:]
	}
	ctx := context.Background()

	//set up user query
	usrQuery := elastic.NewBoolQuery()
	usrQuery = usrQuery.Must(elastic.NewTermQuery("PosterID", strings.ToLower(docID)))
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
	if !(err == io.EOF && res != nil) && err != nil {
		if err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		return "", arrResults, 0, err
	}

	//fmt.Println(res.Hits.TotalHits)

	var report error
	for _, hit := range res.Hits.Hits {
		head, err := uses.ConvertEntryToJournalEntry(eclient, hit.Id, viewerID, true)
		arrResults = append(arrResults, head)
		if err != nil && err != errors.New("This entry is not visible") {
			report = errors.New("One or more problems loading journal entries")

		}
	}

	return res.ScrollId, arrResults, int(res.Hits.TotalHits), report
}
