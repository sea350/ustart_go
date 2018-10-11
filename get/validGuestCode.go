package get

import (
	"context"

	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ValidGuestCode ... Check if guest code is valid
func ValidGuestCode(eclient *elastic.Client, guestCode string) (bool, error) {
	ctx := context.Background()

	termQuery := elastic.NewTermQuery("Code", guestCode)

	searchResult, err := eclient.Search().
		Index(globals.GuestCodeIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return false, err
	}

	exists := searchResult.TotalHits() > 0

	return exists, err

}
