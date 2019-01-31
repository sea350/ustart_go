package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

/*TODO: Make this function much better*/

//IndexBadge ...
// adds a new user document to the ES cluster
// returns err,string. nil, newID if successful.
func IndexBadge(eclient *elastic.Client, newBadge types.Badge) (string, error) {
	// Check if the index exists
	ctx := context.Background()
	var ID string
	exists, err := eclient.IndexExists(globals.BadgeIndex).Do(ctx)
	if err != nil {
		return ID, err
	}
	// If the index doesn't exist, create it and return error.
	if !exists {
		createIndex, Err := eclient.CreateIndex(globals.BadgeIndex).BodyString(globals.MappingBadge).Do(ctx)
		if Err != nil {
			_, _ = eclient.IndexExists(globals.BadgeIndex).Do(ctx)
			panic(Err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
		}

		// Return an error saying it doesn't exist
		return ID, errors.New("Index does not exist")
	}

	// Index the document.
	_, Err := eclient.Index().
		Index(globals.BadgeIndex).
		Type(globals.BadgeType).
		BodyJson(newBadge).
		Id(newBadge.Id).
		Do(ctx)

	if Err != nil {
		return ID, Err
	}

	return newBadge.Id, nil
}
