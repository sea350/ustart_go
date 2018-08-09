package uses

import (
	"time"

	postEntry "github.com/sea350/ustart_go/post/entry"
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectCreatesEntry ... creates a new entry for projects and handles logic/parallel arrays
func ProjectCreatesEntry(eclient *elastic.Client, projID string, posterID string, newContent []rune) (string, error) {
	createdEntry := types.Entry{}
	createdEntry.PosterID = posterID
	createdEntry.Classification = 3
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true
	createdEntry.ReferenceID = projID
	//usr, err := get.GetUserByID(eclient,userID)

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return entryID, err
	}

	err = post.AppendEntryID(eclient, projID, entryID)

	return entryID, err

}

//ProjectCreatesReply ... creates a new reply entry for projects and handles logic/parallel arrays
func ProjectCreatesReply(eclient *elastic.Client, projID string, replyID string, posterID string, newContent []rune) error {
	createdEntry := types.Entry{}
	createdEntry.PosterID = posterID
	createdEntry.Classification = 4
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true
	createdEntry.ReferenceID = projID

	//usr, err := get.GetUserByID(eclient,userID)

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return err
	}

	err = post.AppendEntryID(eclient, projID, entryID)
	if err != nil {
		return err
	}

	err = postEntry.AppendReplyID(eclient, entryID, replyID)

	return err

}

//ProjectCreatesShare ... creates a new share entry for projects and handles logic/parallel arrays
func ProjectCreatesShare(eclient *elastic.Client, projID string, replyID string, posterID string, newContent []rune) error {
	createdEntry := types.Entry{}
	createdEntry.PosterID = posterID
	createdEntry.Classification = 5
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true
	createdEntry.ReferenceID = projID

	//usr, err := get.GetUserByID(eclient,userID)

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return err
	}

	err = post.AppendEntryID(eclient, projID, entryID)
	if err != nil {
		return err
	}

	err = postEntry.AppendShareID(eclient, entryID, replyID)

	return err

}
