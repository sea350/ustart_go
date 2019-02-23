package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//EmailInUse ... checks if email is in use
func EmailInUse(eclient *elastic.Client, theEmail string) (bool, error) {
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("Email", strings.ToLower(theEmail))
	searchResult, err := eclient.Search().
		Index(globals.UserIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return false, err
	} //email might not be in use, but it's still an error

	exists := searchResult.TotalHits() > 0

	return exists, err

}
