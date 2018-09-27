package get

import (
	"context"
	"log"

	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EntryIDsByProjectID ... Returns a list of all entry IDs that a project has
//Takes in the project ID
func EntryIDsByProjectID(eclient *elastic.Client, projectID string) ([]string, error) {
	ctx := context.Background()

	newMatchQuery := elastic.NewMatchQuery("ReferenceID", projectID)

	searchResult, err := eclient.Search().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Query(newMatchQuery).
		Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	var arrayEntryIDs []string
	for _, element := range searchResult.Hits.Hits {
		arrayEntryIDs = append(arrayEntryIDs, element.Id)
	}

	return arrayEntryIDs, err
}
