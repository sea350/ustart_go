package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	entryPost "github.com/sea350/ustart_go/post/entry"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendEntryID ... appends a created entry ID to project
func AppendEntryID(eclient *elastic.Client, usrID string, entryID string) error {
	ctx := context.Background()

	entryPost.EntryLock.Lock()

	usr, err := get.ProjectByID(eclient, usrID)

	if err != nil {
		return errors.New("Project does not exist")
	}

	usr.EntryIDs = append(usr.EntryIDs, entryID)

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(usrID).
		Doc(map[string]interface{}{"EntryIDs": usr.EntryIDs}).
		Do(ctx)

	defer entryPost.EntryLock.Unlock()
	return err

}
