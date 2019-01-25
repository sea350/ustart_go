package properloading

import (
	"context"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ScrollPageUser ...
//Scrolls through docs being loaded on the user wall
func ScrollPageUser(eclient *elastic.Client, docID string, viewerID string, scrollID string) (string, []types.JournalEntry, int, error) {

	ctx := context.Background()

	//set up user query
	usrQuery := elastic.NewBoolQuery()
	usrQuery = usrQuery.Must(elastic.NewTermQuery("PosterID", strings.ToLower(docID)))
	usrQuery = usrQuery.Must(elastic.NewTermsQuery("Classification", 0, 2))
	usrQuery = usrQuery.Must(elastic.NewTermQuery("Visible", true))
	usrQuery = usrQuery.Must(elastic.NewTermQuery("Status", true))

	var arrResults []types.JournalEntry

	scroll := eclient.Scroll().
		Index(globals.EntryIndex).
		Query(usrQuery).
		Sort("TimeStamp", false).
		Size(12)

	scroll = scroll.ScrollId(scrollID)

	res, err := scroll.Do(ctx)
	if err != nil {
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
