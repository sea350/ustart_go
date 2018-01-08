package uses

import (
	"time"

	post "github.com/sea350/ustart_go/post/user"
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

	entryID, err := post.IndexEntry(eclient, createdEntry)
	if err != nil {
		return entryID, err
	}
	err = post.AppendEntryID(eclient, userID, entryID)

	return entryID, err

}
