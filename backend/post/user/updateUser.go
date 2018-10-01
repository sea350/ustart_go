package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/backend/get/user"
	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateUser ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
func UpdateUser(eclient *elastic.Client, userID string, field string, newContent interface{}) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.UserIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.UserByID(eclient, userID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Id(userID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	return err
}
