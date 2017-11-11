package post

import (
	"context"
	"errors"
	"sync"
	"time"

	get "github.com/sea350/ustart_go/get"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

const entryIndex = "test-entry_data"
const entryType = "ENTRY"

var genericEntryUpdateLock sync.Mutex
var likeArrayLock sync.Mutex
var shareArrayLock sync.Mutex
var replyArrayLock sync.Mutex

//IndexEntry ... ADDS NEW ENTRY TO ES RECORDS
//Needs a type entry struct
//returns the new entry's id and an error
func IndexEntry(eclient *elastic.Client, newEntry types.Entry) (string, error) {

	ctx := context.Background()
	var entryID string

	idx, Err := eclient.Index().
		Index(entryIndex).
		Type(entryType).
		BodyJson(newEntry).
		Do(ctx)

	if Err != nil {
		return entryID, Err
	}
	entryID = idx.Id

	return entryID, nil
}

//ReindexEntry ... REPLACES EXISTING ES DOC
//Specify the docid to be replaced and a type Entry struct
//returns an error
func ReindexEntry(eclient *elastic.Client, oldEntry types.Entry, entryID string) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(entryIndex).Do(ctx)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		BodyJson(oldEntry).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}

//UpdateEntry ... UPDATES A SINGLE FEILD IN AN EXISTING ES DOC
//Requires the docID, feild to be modified, and the new content
//Returns an error
func UpdateEntry(eclient *elastic.Client, entryID string, field string, newContent interface{}) error {
	ctx := context.Background()
	//stringified := string(newContent)

	exists, err := eclient.IndexExists(entryIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	genericEntryUpdateLock.Lock()
	defer genericEntryUpdateLock.Unlock()

	_, err = get.GetEntryByID(eclient, entryID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	if err != nil {
		return err
	}
	return nil
}

//AppendLike ... APPENDS A NEW LIKE TO AN EXISTING ENTRY DOC
//Requires the entry docID and the docID of the user that liked
//Returns an error
func AppendLike(eclient *elastic.Client, entryID string, likerID string) error {
	ctx := context.Background()

	likeArrayLock.Lock()
	defer likeArrayLock.Unlock()

	anEntry, err := get.GetEntryByID(eclient, entryID)
	if err != nil {
		return nil
	}

	newLike := types.Like{}
	newLike.UserID = likerID
	newLike.TimeStamp = time.Now()
	anEntry.Likes = append(anEntry.Likes, newLike)
	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"Likes": anEntry.Likes}).
		Do(ctx)

	return err

}

//DeleteLike ... REMOVES A SPECIFIC LIKE FROM AN ENTRY'S Likes
//Requires entry docID and the unliker's docID
//Returns an error
func DeleteLike(eclient *elastic.Client, entryID string, likerID string) error {
	ctx := context.Background()

	likeArrayLock.Lock()
	defer likeArrayLock.Unlock()

	anEntry, err := get.GetEntryByID(eclient, entryID)

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
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"Likes": anEntry.Likes}).
		Do(ctx)

	return err
}

//AppendShareID ... APPENDS A NEW SHARE TO AN EXISTING ENTRY DOC
//Requires the shared entry docID and the docID of the new post
//Returns an error
func AppendShareID(eclient *elastic.Client, entryID string, shareID string) error {
	ctx := context.Background()

	shareArrayLock.Lock()
	defer shareArrayLock.Unlock()

	anEntry, err := get.GetEntryByID(eclient, entryID)
	if err != nil {
		return err
	}
	anEntry.ShareIDs = append(anEntry.ShareIDs, shareID)

	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"ShareIDs": anEntry.ShareIDs}).
		Do(ctx)

	return err

}

//DeleteShareID ... REMOVES A SPECIFIC share FROM AN ENTRY'S Likes
//Requires original entry docID and the share entry's docID
//Returns an error
func DeleteShareID(eclient *elastic.Client, entryID string, shareID string) error {
	ctx := context.Background()

	shareArrayLock.Lock()
	defer shareArrayLock.Unlock()

	anEntry, err := get.GetEntryByID(eclient, entryID)
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
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"ShareIDs": anEntry.ShareIDs}).
		Do(ctx)

	return err

}

//AppendReplyID ... APPENDS A NEW REPLY TO AN EXISTING ENTRY DOC
//Requires the shared entry docID and the docID of the new post
//Returns an error
func AppendReplyID(eclient *elastic.Client, entryID string, replyID string) error {
	ctx := context.Background()

	replyArrayLock.Lock()
	defer replyArrayLock.Unlock()

	anEntry, err := get.GetEntryByID(eclient, entryID)
	if err != nil {
		return err
	}
	anEntry.ShareIDs = append(anEntry.ReplyIDs, replyID)

	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"ReplyIDs": anEntry.ReplyIDs}).
		Do(ctx)

	return err

}

//DeleteReplyID ... REMOVES A SPECIFIC share FROM AN ENTRY'S Likes
//Requires original entry docID and the share entry's docID
//Returns an error
func DeleteReplyID(eclient *elastic.Client, entryID string, replyID string) error {
	ctx := context.Background()

	replyArrayLock.Lock()
	defer replyArrayLock.Unlock()

	anEntry, err := get.GetEntryByID(eclient, entryID)
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

	anEntry.ShareIDs = append(anEntry.ReplyIDs[:idx], anEntry.ReplyIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"ReplyIDs": anEntry.ReplyIDs}).
		Do(ctx)

	return err
}
