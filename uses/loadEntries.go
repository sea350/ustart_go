package uses

import (
import "fmt"
import "fmt"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//LoadEntries ... Loads a list of entries as journal entries, if an entry is invisible it is skipped
//Requires an array of entry ids
//Returns an of the data for those ids as journal entries, and an error
func LoadEntries(eclient *elastic.Client, loadList []string) ([]types.JournalEntry, error) {

	var entries []types.JournalEntry

	for _, entryID := range loadList {
		jEntry, err := ConvertEntryToJournalEntry(eclient, entryID)
		if err != nil {
			return entries, err
		}

		if !jEntry.Element.Visible {
			continue
		}

		entries = append(entries, jEntry)
	}

	
	var noClassOnes []types.JournalEntry
	for _, entry := range entries {
		if entry.Element.Classification != 1 {
			fmt.Println(entry.
			noClassOnes = append(noClassOnes, entry)

		}
	}
	return noClassOnes, nil
}
