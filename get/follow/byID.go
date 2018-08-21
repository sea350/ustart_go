package get

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
	//post "github.com/sea350/ustart_go/post"
)

//ByID ...
func ByID(eclient *elastic.Client, userID string) (string, types.Follow, error) {
	ctx := context.Background() //intialize context background
	var foll types.Follow       //initialize follow
	var follID string           //initialize follow ID
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("DocID", strings.ToLower(userID)))
	searchResult, err := eclient.Search(). //Get returns doc type, index, etc.
						Index(globals.FollowIndex).
						Type(globals.FollowType).
						Do(ctx)

	if err != nil {
		return "", foll, err
	}

	if searchResult.Hits.TotalHits > 2 {
		return "", foll, errors.New("More than one result found")
	} else if searchResult.Hits.TotalHits < 1 {
		return "", foll, errors.New("No results found")
	}

	for _, hit := range searchResult.Hits.Hits {
		err = json.Unmarshal(*hit.Source, &foll) //unmarshal type RawMessage into user struct
		follID = hit.Id
	}

	return follID, foll, err

}
