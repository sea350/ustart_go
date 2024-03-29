package uses

import (
	getEntry "github.com/sea350/ustart_go/get/entry"
	getProject "github.com/sea350/ustart_go/get/project"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ProjectPage ...
func ProjectPage(eclient *elastic.Client, projectURL string, viewerID string) (types.Project, []types.JournalEntry, string, int, error) {
	maxPull := 20
	counter := 0
	var proj types.Project

	var entries []types.JournalEntry
	var memberClass int

	projID, err := getProject.ProjectIDByURL(eclient, projectURL)
	if err != nil {
		return proj, entries, projID, memberClass, err
	}

	viewer, err := getProject.ProjectByID(eclient, viewerID)
	if err != nil {
		return proj, entries, projID, memberClass, err
	}

	for _, element := range proj.Members {
		if element.MemberID == viewerID {
			memberClass = element.Role
			break
		} else {
			memberClass = -1
		}
	}

	for _, i := range proj.EntryIDs {
		//goes through the user's entries
		entry, err := getEntry.EntryByID(eclient, i)
		if err != nil {
			return proj, entries, projID, memberClass, err
		}
		if !entry.Visible {
			continue
		} //checks if entry is visible
		//if invisible, then skip

		var newEntry types.JournalEntry
		newEntry.Element = entry
		newEntry.FirstName = viewer.Name
		//newEntry.LastName = viewer.LastName
		newEntry.NumReplies = len(entry.ReplyIDs)
		newEntry.NumLikes = len(entry.Likes)
		newEntry.NumShares = len(entry.ShareIDs)

		//check if invis
		entries = append(entries, newEntry)
		counter++
		if counter > maxPull {
			break
		}
	}

	return proj, entries, projID, counter, nil //not sure of what exactly this returns, beyond entries

}
