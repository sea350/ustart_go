package post

import (
	"context"

	get "github.com/sea350/ustart_go/get/entry"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendShareID ... APPENDS A NEW SHARE TO AN EXISTING ENTRY DOC
//Requires the shared entry docID and the docID of the new post
//Returns an error
func AppendShareID(eclient *elastic.Client, entryID string, shareID string) error {
	ctx := context.Background()

	ShareArrayLock.Lock()

	anEntry, err := get.EntryByID(eclient, entryID)
	if err != nil {
		return err
	}
	anEntry.ShareIDs = append(anEntry.ShareIDs, shareID)

	_, err = eclient.Update().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		Doc(map[string]interface{}{"ShareIDs": anEntry.ShareIDs}).
		Do(ctx)

	defer ShareArrayLock.Unlock()
	return err

}
