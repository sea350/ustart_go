package post

import (
	"context"
	"time"

	getEntry "github.com/sea350/ustart_go/get/entry"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//AppendLike ... APPENDS A NEW LIKE TO AN EXISTING ENTRY DOC
//Requires the entry docID and the docID of the user that liked
//Returns an error
func AppendLike(eclient *elastic.Client, entryID string, likerID string) error {
	ctx := context.Background()

	LikeArrayLock.Lock()

	anEntry, err := getEntry.EntryByID(eclient, entryID)
	if err != nil {
		return nil
	}

	newLike := types.Like{}
	newLike.UserID = likerID
	newLike.TimeStamp = time.Now()
	anEntry.Likes = append(anEntry.Likes, newLike)
	_, err = eclient.Update().
		Index(globals.EntryIndex).
		Type(globals.EntryType).
		Id(entryID).
		Doc(map[string]interface{}{"Likes": anEntry.Likes}).
		Do(ctx)

	defer LikeArrayLock.Unlock()

	return err

}
