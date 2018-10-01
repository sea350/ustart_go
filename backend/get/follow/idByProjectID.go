package get

import (
	"context"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IDByProjectID ... Gets a follow object by project ID
//Takes in a project ID
//Returns the Follow ID
func IDByProjectID(eclient *elastic.Client, projectID string) (string, error) {
	ctx := context.Background() //intialize context background
	var followID string
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("DocID", strings.ToLower(projectID)))
	searchResult, err := eclient.Search(). //Get returns doc type,s index, etc.
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
