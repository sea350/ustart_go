package uses

import (
	getEntry "github.com/sea350/ustart_go/get/entry"
	getUser "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ConvertEntryToJournalEntry ... load all relevant data of a single entry into a journal entry struct
//Requires entry docID
//Returns entry as a type JournalEntry and an error
func ConvertEntryToJournalEntry(eclient *elastic.Client, entryID string, viewerID string, enableRecursion bool) (types.JournalEntry, error) {
	var newJournalEntry types.JournalEntry

	newJournalEntry.ElementID = entryID

	liked, err := IsLiked(eclient, entryID, viewerID)
	if err != nil {
		return newJournalEntry, err
	}

	newJournalEntry.Liked = liked

	entry, err := getEntry.EntryByID(eclient, entryID)
	if err != nil || !entry.Visible {
		return newJournalEntry, err
	}
	newJournalEntry.Element = entry
	newJournalEntry.NumShares = len(entry.ShareIDs)
	newJournalEntry.NumLikes = len(entry.Likes)
	newJournalEntry.NumReplies = len(entry.ReplyIDs)

	usr, err := getUser.UserByID(eclient, entry.PosterID)
	if err != nil {
		return newJournalEntry, err
	}
	newJournalEntry.FirstName = usr.FirstName
	newJournalEntry.LastName = usr.LastName
	newJournalEntry.Image = usr.Avatar
	if entry.Classification == 2 && enableRecursion {
		newJournalEntry.ReferenceElement, err = ConvertEntryToJournalEntry(eclient, entry.ReferenceEntry, viewerID, false)
	}

	return newJournalEntry, err
}
