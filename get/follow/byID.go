package get

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
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
		exists, err := eclient.IndexExists(globals.FollowIndex).Do(ctx)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		// If the index doesn't exist, create it and return error.
		if !exists {
			createIndex, Err := eclient.CreateIndex(globals.FollowIndex).BodyString(globals.MappingFollow).Do(ctx)
			if Err != nil {
				_, _ = eclient.IndexExists(globals.FollowIndex).Do(ctx)
				panic(Err)
			}
			// TODO fix this.
			if !createIndex.Acknowledged {
			}

			// Return an error saying it doesn't exist
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}

		var newFollowing = make(map[string]bool)
		var newFollowers = make(map[string]bool)
		var newBell = make(map[string]bool)
		var newFollow = types.Follow{
			DocID: userID,

			UserFollowers:    newFollowers,
			UserFollowing:    newFollowing,
			ProjectFollowers: newFollowers,
			ProjectFollowing: newFollowing,
			EventFollowers:   newFollowers,
			EventFollowing:   newFollowing,
			UserBell:         newBell,
			ProjectBell:      newBell,
			EventBell:        newBell,
		}
		// Index the document.
		_, Err := eclient.Index().
			Index(globals.FollowIndex).
			Type(globals.FollowType).
			BodyJson(newFollow).
			Do(ctx)
		if Err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(Err)
		}
	}

	for _, hit := range searchResult.Hits.Hits {
		err = json.Unmarshal(*hit.Source, &foll) //unmarshal type RawMessage into user struct
		follID = hit.Id
	}

	return follID, foll, err

}
