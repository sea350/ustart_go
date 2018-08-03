package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IndexImg ...
// adds a new image document to the ES cluster
// returns err, nil if successful.
func IndexImg(eclient *elastic.Client, newImg types.Img) (string, error) {
	// Check if the index exists
	ctx := context.Background()
	var id string
	exists, err := eclient.IndexExists(globals.ImgIndex).Do(ctx)
	if err != nil {
		return id, err
	}
	// If the index doesn't exist, create it and return error.
	if !exists {
		createIndex, Err := eclient.CreateIndex(globals.ImgIndex).Do(ctx)
		if Err != nil {
			_, _ = eclient.IndexExists(globals.ImgIndex).Do(ctx)
			panic(Err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
			// Return an error saying it doesn't exist
			return id, errors.New("Index does not exist")
		}
	}

	// Index the document.
	newI, Err := eclient.Index().
		Index(globals.ImgIndex).
		BodyJson(newImg).
		Do(ctx)

	if Err != nil {
		return id, Err
	}

	return newI.Id, nil
}
