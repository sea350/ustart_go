package get

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//ByID ...
//
func FollowExists(eclient *elastic.Client, userID string) (bool, error) {
	ctx := context.Background() //intialize context background

	query := elastic.NewTermQuery("DocID", userID)
	searchResult, err := eclient.Search(). //Get returns doc type, index, etc.
						Index(globals.FollowIndex).
						Type(globals.FollowType).
						Query(query).
						Do(ctx)

	if err != nil {
		return false, err
	}

	return searchResult.Hits.TotalHits > 0, err

}
