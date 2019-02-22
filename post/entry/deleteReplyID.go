package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/entry"
	globals "github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

//DeleteReplyID ... REMOVES A SPECIFIC share FROM AN ENTRY'S Likes
//Requires original entry docID and the share entry's docID
//Returns an error
func DeleteReplyID(eclient *elastic.Client, entryID string, replyID string) error {
	ctx := context.Background()

	ReplyArrayLock.Lock()
	defer ReplyArrayLock.Unlock()

	anEntry, err := get.EntryByID(eclient, entryID)
	if err != nil {
		return nil
	}

	idx := -1
	for i := range anEntry.ReplyIDs {
		if replyID == anEntry.ReplyIDs[i] {
			idx = i
			break
		}
	}
	if idx == -1 {
		return errors.New("Reply not found")
	}

	anEntry.ReplyIDs = append(anEntry.ReplyIDs[:idx], anEntry.ReplyIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		Doc(map[string]interface{}{"ReplyIDs": anEntry.ReplyIDs}).
		Do(ctx)

	return err
}
