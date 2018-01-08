package uses

import (
	"time"

	entryPost "github.com/sea350/ustart_go/post/entry"
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectCreatesEntry ...
func ProjectCreatesEntry(eclient *elastic.Client, projID string, newContent []rune) error {
	createdEntry := types.Entry{}
	createdEntry.PosterID = projID
	createdEntry.Classification = 0
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true

	//usr, err := get.GetUserByID(eclient,userID)

	entryID, err := entryPost.IndexEntry(eclient, createdEntry)
	if err != nil {
		return err
	}
	err = post.AppendEntryID(eclient, projID, entryID)

	return err

}
