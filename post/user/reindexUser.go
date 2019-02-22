package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ReindexUser ...
//  Add a new user to ES.
//  Returns an error, nil if successful
func ReindexUser(eclient *elastic.Client, userID string, userAcc types.User) error {

	ctx := context.Background()
	exists, err := eclient.IndexExists(globals.UserIndex).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(userID).
		BodyJson(userAcc).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
