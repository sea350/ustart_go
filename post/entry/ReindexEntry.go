package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ReindexEntry ... REPLACES EXISTING ES DOC
//Specify the docid to be replaced and a type Entry struct
//returns an error
func ReindexEntry(eclient *elastic.Client, oldEntry types.Entry, entryID string) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.EntryIndex).Do(ctx)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		BodyJson(oldEntry).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
