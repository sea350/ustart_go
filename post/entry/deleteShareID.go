package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/entry"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteShareID ... REMOVES A SPECIFIC share FROM AN ENTRY'S Likes
//Requires original entry docID and the share entry's docID
//Returns an error
func DeleteShareID(eclient *elastic.Client, entryID string, shareID string) error {
	ctx := context.Background()

	ShareArrayLock.Lock()

	anEntry, err := get.EntryByID(eclient, entryID)
	if err != nil {
		return nil
	}

	idx := -1
	for i := range anEntry.ShareIDs {
		if shareID == anEntry.ShareIDs[i] {
			idx = i
			break
		}
	}
	if idx == -1 {
		return errors.New("Share not found")
	}

	anEntry.ShareIDs = append(anEntry.ShareIDs[:idx], anEntry.ShareIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		Doc(map[string]interface{}{"ShareIDs": anEntry.ShareIDs}).
		Do(ctx)

	defer ShareArrayLock.Unlock()
	return err

}
