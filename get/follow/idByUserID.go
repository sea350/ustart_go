package get

import (
	"context"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
	//post "github.com/sea350/ustart_go/post"
)

//IDByUserID ...  Gets follow ID by user ID
//Requires the user's ID
//Returns the follow ID and error
func IDByUserID(eclient *elastic.Client, userID string) (string, error) {
	ctx := context.Background() //intialize context background
	var followID string
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("DocID", strings.ToLower(userID)))
	searchResult, err := eclient.Search(). //Get returns doc type, index, etc.
						Index(globals.FollowIndex).
						Type(globals.FollowType).
						Do(ctx)

	if err != nil {
		return followID, err
	}

	if searchResult.Hits.TotalHits > 1 {
		return followID, errors.New("More than one result found")
	} else if searchResult.Hits.TotalHits < 1 {
		return followID, errors.New("No results found")
	}

	for _, hit := range searchResult.Hits.Hits {
		followID = hit.Id
	}

	return followID, err

}
