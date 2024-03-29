package post

import (
	get "github.com/sea350/ustart_go/get/entry"
	//post "github.com/sea350/ustart_go/post/entry"
	elastic "github.com/olivere/elastic"
)

//AppendReplyID ... APPENDS A NEW REPLY TO AN EXISTING ENTRY DOC
//Requires the shared entry docID and the docID of the new post
//Returns an error
func AppendReplyID(eclient *elastic.Client, entryID string, replyID string) error {

	ReplyArrayLock.Lock()

	anEntry, err := get.EntryByID(eclient, entryID)
	if err != nil {
		return err
	}
	anEntry.ReplyIDs = append(anEntry.ReplyIDs, replyID)

	err = UpdateEntry(eclient, entryID, "ReplyIDs", anEntry.ReplyIDs)

	defer ReplyArrayLock.Unlock()
	return err

}
