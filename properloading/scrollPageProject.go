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

//ScrollPageProject ...
//Scrolls through docs being loaded on project page
func ScrollPageProject(eclient *elastic.Client, docID string, viewerID string, scrollID string) (string, []types.JournalEntry, int, error) {

	ctx := context.Background()

	var dash byte = '-'
	var underscore byte = '_'
	for docID[0] == dash || docID[0] == underscore {
		docID = docID[1:]
	}
	for docID[len(docID)-1] == dash || docID[len(docID)-1] == underscore {
		docID = docID[:len(docID)-1]
	}

	//set up project query
	projQuery := elastic.NewBoolQuery()
	projQuery = projQuery.Must(elastic.NewTermQuery("ReferenceID", strings.ToLower(docID)))
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
	if !(err == io.EOF && res != nil) && err != nil {
		if err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
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
