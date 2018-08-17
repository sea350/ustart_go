package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

/*TODO: Make this function much better*/

//ReindexFollow ...
// reindexes follow document to the ES cluster
// returns  nil,  if successful.
func ReindexFollow(eclient *elastic.Client, followID string, follow types.Follow) error {
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

	// Index the document.
	_, Err := eclient.Index().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		BodyJson(follow).
		Id(followID).
		Do(ctx)

	return Err
}
