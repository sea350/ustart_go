package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateGuestCode ... Updates specified field of GuestCode
func UpdateGuestCode(eclient *elastic.Client, codeID string, field string, newContent interface{}) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.GuestCodeIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Update().Index(globals.GuestCodeIndex).
		Type(globals.GuestCodeType).
		Id(codeID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)
	return err
}
