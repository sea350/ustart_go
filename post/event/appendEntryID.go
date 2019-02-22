package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	globals "github.com/sea350/ustart_go/globals"
	entryPost "github.com/sea350/ustart_go/post/entry"
	elastic "github.com/olivere/elastic"
)

//AppendEntryID ... appends a created entry ID to user
func AppendEntryID(eclient *elastic.Client, usrID string, entryID string) error {
	ctx := context.Background()
	entryPost.EntryLock.Lock()

	evnt, err := get.EventByID(eclient, usrID)

	if err != nil {
		return errors.New("Event does not exist")
	}

	evnt.EntryIDs = append(evnt.EntryIDs, entryID)

	_, err = eclient.Update().
		Index(globals.EventIndex).
		Type(globals.EventType).
		Id(usrID).
		Doc(map[string]interface{}{"EntryIDs": evnt.EntryIDs}).
		Do(ctx)

	defer entryPost.EntryLock.Unlock()
	return err

}
