package delete

import (
	elastic "github.com/olivere/elastic"
	//"errors"

	"context"

	globals "github.com/sea350/ustart_go/globals"
	//"golang.org/x/crypto/bcrypt"
)

//Entry ...
//
func Entry(eclient *elastic.Client, entryID string) error {

	ctx := context.Background()

	//delete the widget from ES
	_, err := eclient.Delete().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		Do(ctx)

	return err

}
