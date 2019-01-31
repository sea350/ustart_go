package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ReindexBadge ...
//  Add a new badge to ES.
//  Returns an error, nil if successful
func ReindexBadge(eclient *elastic.Client, badgeID string, theBadge types.Badge) error {

	ctx := context.Background()
	exists, err := eclient.IndexExists(globals.BadgeIndex).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.BadgeIndex).
		Type(globals.BadgeType).
		Id(badgeID).
		BodyJson(theBadge).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
