package post

import (
	"context"

	get "github.com/sea350/ustart_go/get/entry"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendReplyID ... APPENDS A NEW REPLY TO AN EXISTING ENTRY DOC
//Requires the shared entry docID and the docID of the new post
//Returns an error
func AppendReplyID(eclient *elastic.Client, entryID string, replyID string) error {
	ctx := context.Background()

	replyArrayLock.Lock()
	defer replyArrayLock.Unlock()

	anEntry, err := get.EntryByID(eclient, entryID)
	if err != nil {
		return err
	}
	anEntry.ShareIDs = append(anEntry.ReplyIDs, replyID)

	_, err = eclient.Update().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		Doc(map[string]interface{}{"ReplyIDs": anEntry.ReplyIDs}).
		Do(ctx)

	return err

}
