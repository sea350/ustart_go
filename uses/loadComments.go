package uses

import (
	"errors"
	"log"

	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//LoadComments ... Loads the replies to a specific entry limited by limits
//NOTE set uppper bound to -1 to pull to the end of the array
//Returns the parent entry as a JournalEntry, an array of replies, and an error
//NOTE, if the entry is set to invisible it is skipped
func LoadComments(eclient *elastic.Client, entryID string, viewerID string, lowerBound int, upperBound int) (types.JournalEntry, []types.JournalEntry, error) {
	var entries []types.JournalEntry
	var parent types.JournalEntry
	var start int
	var finish int

	if lowerBound < 0 {
		return parent, entries, errors.New("Lower Bound limit is out of bounds")
	}

	parent, err := ConvertEntryToJournalEntry(eclient, entryID, viewerID, true)
	if err != nil && err != errors.New("This entry is not visible") {
		return parent, entries, err
	}
	if upperBound == -1 {
		finish = 0
	} else if len(parent.Element.ReplyIDs)-upperBound < 0 {
		finish = 0
	} else {
		finish = len(parent.Element.ReplyIDs) - upperBound
	}

	start = (len(parent.Element.ReplyIDs) - 1) - lowerBound
	for i := start; i >= finish; i-- {
		jEntry, err := ConvertEntryToJournalEntry(eclient, parent.Element.ReplyIDs[i], viewerID, true)
		if err != nil && err != errors.New("This entry is not visible") {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}

		if !jEntry.Element.Visible && finish > 0 {
			finish--
			continue
		}

		entries = append(entries, jEntry)
	}

	return parent, entries, err
}
