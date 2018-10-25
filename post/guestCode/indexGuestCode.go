package post

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	"gopkg.in/olivere/elastic.v5"
)

//IndexGuestCode ... Adds a new guest code document and checks if a GuestCode index already exists
//Needs a GuestCode struct
//Returns ID of new GuestCode and an error if it exists
func IndexGuestCode(eclient *elastic.Client, guestCode types.GuestCode) (string, error) {
	var ID string
	//Checks if index already exists
	ctx := context.Background()
	exists, err := eclient.IndexExists(globals.GuestCodeIndex).Do(ctx)
	if err != nil {
		return ID, err
	}

	//If index does not exist, create a new one
	if !exists {
		createIndex, err := eclient.CreateIndex(globals.GuestCodeIndex).Do(ctx)
		if err != nil {
			return ID, err
		}
		if !createIndex.Acknowledged {
			panic(err)
		}
	}

	//Index document
	newGuestCode, err := eclient.Index().
		Index(globals.GuestCodeIndex).
		Type(globals.GuestCodeType).
		BodyJson(guestCode).
		Do(ctx)
	if err != nil {
		return ID, err
	}

	return newGuestCode.Id, nil
}
