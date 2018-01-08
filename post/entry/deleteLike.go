package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/entry"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteLike ... REMOVES A SPECIFIC LIKE FROM AN ENTRY'S Likes
//Requires entry docID and the unliker's docID
//Returns an error
func DeleteLike(eclient *elastic.Client, entryID string, likerID string) error {
	ctx := context.Background()

	LikeArrayLock.Lock()

	anEntry, err := get.EntryByID(eclient, entryID)

	idx := -1
	for i := range anEntry.Likes {
		if likerID == anEntry.Likes[i].UserID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return errors.New("Like not found")
	}

	anEntry.Likes = append(anEntry.Likes[:idx], anEntry.Likes[idx+1:]...)

	_, err = eclient.Update().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		Doc(map[string]interface{}{"Likes": anEntry.Likes}).
		Do(ctx)

	defer LikeArrayLock.Unlock()
	return err
}
