package post

import (
	"context"
	"errors"

	//get "github.com/sea350/ustart_go/get/entry"
	get "github.com/sea350/ustart_go/get/entry"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateEntry ... UPDATES A SINGLE FEILD IN AN EXISTING ES DOC
//Requires the docID, feild to be modified, and the new content
//Returns an error
func UpdateEntry(eclient *elastic.Client, entryID string, field string, newContent interface{}) error {
	ctx := context.Background()
	//stringified := string(newContent)

	exists, err := eclient.IndexExists(globals.EntryIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	GenericEntryUpdateLock.Lock()

	_, err = get.EntryByID(eclient, entryID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	defer GenericEntryUpdateLock.Unlock()

	return err
}

func UpdateEditEntry(eclient *elastic.Client, entryID string, field string, newContent interface{}) error {
	ctx := context.Background()
	//stringified := string(newContent)

	exists, err := eclient.IndexExists(globals.EntryIndex).Do(ctx)

	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	GenericEntryUpdateLock.Lock()

	_, err = get.EntryByID(eclient, entryID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	defer GenericEntryUpdateLock.Unlock()

	return err
}
