package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/backend/globals"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

/*TODO: Make this function much better*/

//IndexFollow ...
// adds a new follow document to the ES cluster
// returns err,string. nil, newID if successful.
func IndexFollow(eclient *elastic.Client, userID string) error {
	// Check if the index exists
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.FollowIndex).Do(ctx)
	if err != nil {
		return err
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
		return errors.New("Index does not exist")
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

	return Err
}
