package post

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//IndexEntry ... ADDS NEW ENTRY TO ES RECORDS
//Needs a type entry struct
//returns the new entry's id and an error
func IndexEntry(eclient *elastic.Client, newEntry types.Entry) (string, error) {

	ctx := context.Background()
	var entryID string

	idx, Err := eclient.Index().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		BodyJson(newEntry).
		Do(ctx)

	if Err != nil {
		return entryID, Err
	}
	entryID = idx.Id

	return entryID, nil
}
