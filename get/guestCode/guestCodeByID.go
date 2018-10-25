package get

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//GuestCodeByID ... Gets a guestCode by it's ID
func GuestCodeByID(eclient *elastic.Client, codeID string) (types.GuestCode, error) {
	ctx := context.Background()
	var guestCode types.GuestCode

	newQuery := elastic.NewTermsQuery("Code", strings.ToLower(codeID))
	searchResults, err := eclient.Search().
		Index(globals.GuestCodeIndex).
		Query(newQuery).
		Do(ctx)
	if err != nil {
		return guestCode, err
	}

	if searchResults.Hits.TotalHits != 1 {
		log.Println(searchResults.Hits.Hits, " results")
		return guestCode, errors.New("Not 1 result")
	}
	for _, hit := range searchResults.Hits.Hits {
		err = json.Unmarshal(*hit.Source, &guestCode)
		if err != nil {
			return guestCode, err
		}
		break
	}

	return guestCode, err
}
