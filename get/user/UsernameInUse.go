package get

import (
	"context"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//UsernameInUse ... checks if username is in use
func UsernameInUse(eclient *elastic.Client, username string) (bool, error) {
	ctx := context.Background()
	usernameLower := strings.ToLower(username)
	termQuery := elastic.NewTermQuery("Username", usernameLower)

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
