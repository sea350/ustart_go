package uses

import (
	"time"

	postEntry "github.com/sea350/ustart_go/post/entry"
	post "github.com/sea350/ustart_go/post/event"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventCreatesEntry ... creates a new entry for events and handles logic/parallel arrays
func EventCreatesEntry(eclient *elastic.Client, eventID string, posterID string, newContent []rune) (string, error) {
	createdEntry := types.Entry{}
	createdEntry.PosterID = posterID
	createdEntry.Classification = 0
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true
	createdEntry.ReferenceID = eventID

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return entryID, err
	}

	err = post.AppendEntryID(eclient, eventID, entryID)
	return entryID, err
}

//EventCreatesReply ... creates a new reply entry for events and handles logic/parallel arrays
func EventCreatesReply(eclient *elastic.Client, eventID string, replyID string, posterID string, newContent []rune) error {
	createdEntry := types.Entry{}
	createdEntry.PosterID = posterID
	createdEntry.Classification = 1
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true
	createdEntry.ReferenceID = eventID

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return err
	}

	err = post.AppendEntryID(eclient, eventID, entryID)
	if err != nil {
		return err
	}

	err = postEntry.AppendReplyID(eclient, entryID, replyID)
	return err
}

//EventCreatesShare ... creates a new share entry for events and handles logic/parallel arrays
func EventCreatesShare(eclient *elastic.Client, eventID string, replyID string, posterID string, newContent []rune) error {
	createdEntry := types.Entry{}
	createdEntry.PosterID = posterID
	createdEntry.Classification = 2
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true
	createdEntry.ReferenceID = eventID

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return err
	}

	err = post.AppendEntryID(eclient, eventID, entryID)
	if err != nil {
		return err
	}

	err = postEntry.AppendShareID(eclient, entryID, replyID)

	return err
}
