package get

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UsernameInUse ... checks if username is in use
func UsernameInUse(eclient *elastic.Client, username string) (bool, error) {
	ctx := context.Background()

	termQuery := elastic.NewTermQuery("Username", username)
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
