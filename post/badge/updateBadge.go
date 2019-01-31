package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/badge"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateBadge...
//  Change a single field of the ES Document
//  Return an error, nil if successful
func UpdateBadge(eclient *elastic.Client, badgeID string, field string, newContent interface{}) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.BadgeIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.BadgeByID(eclient, badgeID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.BadgeIndex).
		Type(globals.BadgeType).
		Id(badgeID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	return err
}
