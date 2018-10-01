package get

import (
	"context"
	"encoding/json"
	"log"

	globals "github.com/sea350/ustart_go/backend/globals"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ByID ...
//
func ByID(eclient *elastic.Client, userID string) (string, types.Follow, error) {
	ctx := context.Background() //intialize context background
	var foll types.Follow       //initialize follow
	var follID string           //initialize follow ID

	query := elastic.NewTermQuery("DocID", userID)
	searchResult, err := eclient.Search(). //Get returns doc type, index, etc.
						Index(globals.FollowIndex).
						Type(globals.FollowType).
						Query(query).
						Do(ctx)

	if err != nil {
		return "", foll, err
	}

	if searchResult.Hits.TotalHits > 1 {
		for idx, hit := range searchResult.Hits.Hits {
			if idx < int(searchResult.Hits.TotalHits)-1 {
				_, err = eclient.Delete().
					Index(globals.FollowIndex).
					Type(globals.FollowType).
					Id(hit.Id).
					Do(ctx)
				if err != nil {
					return "", foll, err
				}
			} else {
				break
			}
		}

	} else if searchResult.Hits.TotalHits == 0 {
		// var newFollowing = make(map[string]bool)
		// var newFollowers = make(map[string]bool)
		// var newBell = make(map[string]bool)
		var newFollow = types.Follow{
			DocID: userID,
		}
		// UserFollowers:    newFollowers,
		// UserFollowing:    newFollowing,
		// ProjectFollowers: newFollowers,
		// ProjectFollowing: newFollowing,
		// EventFollowers:   newFollowers,
		// EventFollowing:   newFollowing,
		// UserBell:         newBell,
		// ProjectBell:      newBell,
		// EventBell:        newBell,

		// Index the document.
		newDoc, Err := eclient.Index().
			Index(globals.FollowIndex).
			Type(globals.FollowType).
			BodyJson(newFollow).
			Do(ctx)
		if Err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(Err)
		}

		return newDoc.Id, newFollow, err
	}

	for _, hit := range searchResult.Hits.Hits {
		err = json.Unmarshal(*hit.Source, &foll) //unmarshal type RawMessage into user struct
		follID = hit.Id
	}

	if len(foll.UserFollowers) == 0 {
		foll.UserFollowers = make(map[string]bool)
	}

	if len(foll.UserFollowing) == 0 {
		foll.UserFollowing = make(map[string]bool)
	}

	if len(foll.ProjectFollowers) == 0 {
		foll.ProjectFollowers = make(map[string]bool)
	}

	return follID, foll, err

}
