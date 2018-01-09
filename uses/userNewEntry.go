package uses

import (
	"time"

	postEntry "github.com/sea350/ustart_go/post/entry"
	postUser "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UserNewEntry ... CREATES AN ORIGINAL POST FROM A USER
//Requires the user's docID and the content of the post
//Returns a the new entry's docID, and an error
func UserNewEntry(eclient *elastic.Client, userID string, newContent []rune) (string, error) {
	var entryID string

	createdEntry := types.Entry{}
	createdEntry.PosterID = userID
	createdEntry.Classification = 0
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return entryID, err
	}
	err = postUser.AppendEntryID(eclient, userID, entryID)

	return entryID, err

}
