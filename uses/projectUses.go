package uses

import (
	"errors"
	"time"

	get "github.com/sea350/ustart_go/get"
	post "github.com/sea350/ustart_go/post"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

func CreateProject(eclient *elastic.Client, title string, description []rune, makerID string) error {
	var newProj types.Project
	newProj.Name = title
	newProj.Description = description
	newProj.Visible = true
	newProj.CreationDate = time.Now()

	var maker types.Member
	maker.JoinDate = time.Now()
	maker.MemberID = makerID
	maker.Role = 0
	maker.Title = "Creator"
	maker.Visible = true

	newProj.Members = append(newProj.Members, maker)

	id, err := post.IndexProject(eclient, newProj)
	if err != nil {
		return err
	}
	var addProj types.ProjectInfo
	addProj.ProjectID = id
	addProj.Visible = true
	err = post.AppendToUser(eclient, makerID, "Projects", addProj)

	return err

}

func ChangeProjectNameAndDescription(eclient *elastic.Client, projectID string, newName string, newDescription []rune) error {
	err := post.UpdateProject(eclient, projectID, "Name", newName)
	if err != nil {
		return err
	}
	err = post.UpdateProject(eclient, projectID, "Description", newDescription)
	return err
}

func ChangeProjectLocation(eclient *elastic.Client, projectID string, country string, state string, city string, zip string) error {
	var newLoc types.LocStruct
	newLoc.Country = country
	newLoc.State = state
	newLoc.City = city
	newLoc.Zip = zip

	err := post.UpdateProject(eclient, projectID, "Location", newLoc)
	return err

}

func ChangeProjectCategory(eclient *elastic.Client, projectID string, category string) error {
	err := post.UpdateProject(eclient, projectID, "Category", category)
	return err
}

func ChangeProjectURL(eclient *elastic.Client, projectID string, newURL string) error {
	_, err := get.GetProjectByURL(eclient, newURL)
	//if (err != nil){ return err}
	if err != nil {
		return errors.New("That url is already taken")
	}
	err = post.UpdateProject(eclient, projectID, "URLName", newURL)
	return err
}

func ManageMembers(eclient *elastic.Client, projectID string, newMemberConfig []types.Member) error {
	err := post.UpdateProject(eclient, projectID, "Members", newMemberConfig)
	return err
}

func RemoveMember(eclient *elastic.Client, projectID string, userID string) error {
	usr, err := get.GetUserByID(eclient, userID)
	if err != nil {
		return err
	}
	proj, projErr := get.GetProjectByID(eclient, projectID)
	if err != nil {
		return projErr
	}

	var usrIdx, projIdx int
	var infoProj types.ProjectInfo
	for idx := range usr.Projects {
		if usr.Projects[idx].ProjectID == projectID {
			usrIdx = idx
			infoProj = usr.Projects[idx]
			break
		}
	}

	usrErr := post.RemoveFromUser(eclient, userID, "Projects", usrIdx, infoProj)
	if usrErr != nil {
		return usrErr
	}

	for index := range proj.Members {
		if proj.Members[index].MemberID == projectID {
			projIdx = index
			break
		}
	}

	projErr = post.DeleteMember(eclient, projectID, proj.Members[projIdx])
	if projErr != nil {
		return projErr
	}

	return nil

}

func RequestMember(eclient *elastic.Client, projectID string, userID string) error {
	err := post.AppendToUser(eclient, userID, "ReceivedProjReq", projectID)
	if err != nil {
		return err
	}
	err = post.AppendMemberReqSent(eclient, projectID, userID)
	return err
}

func ProjectPage(eclient *elastic.Client, projectURL string, viewerID string) (types.Project, []types.JournalEntry, string, int, error) {
	maxPull := 20
	counter := 0
	var proj types.Project

	var entries []types.JournalEntry
	var memberClass int

	projID, err := get.GetProjectIDByURL(eclient, projectURL)
	if err != nil {
		return proj, entries, projID, memberClass, err
	}

	viewer, err := get.GetUserByID(eclient, viewerID)
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
		entry, err := get.GetEntryByID(eclient, i)
		if err != nil {
			return proj, entries, projID, memberClass, err
		}
		if !entry.Visible {
			continue
		} //checks if entry is visible
		//if invisible, then skip

		var newEntry types.JournalEntry
		newEntry.Element = entry
		newEntry.FirstName = viewer.FirstName
		newEntry.LastName = viewer.LastName
		newEntry.NumReplies = len(entry.ReplyIDs)
		newEntry.NumLikes = len(entry.Likes)
		newEntry.NumShares = len(entry.ShareIDs)

		//check if invis
		entries = append(entries, newEntry)
		counter += 1
		if counter > maxPull {
			break
		}
	}

	return proj, entries, projID, counter, nil //not sure of what exactly this returns, beyond entries

}

func ProjectCreatesEntry(eclient *elastic.Client, projID string, newContent []rune) error {
	createdEntry := types.Entry{}
	createdEntry.PosterID = projID
	createdEntry.Classification = 0
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true

	//usr, err := get.GetUserByID(eclient,userID)

	entryID, err := post.IndexEntry(eclient, createdEntry)
	if err != nil {
		return err
	}
	err = post.AppendEntryID(eclient, projID, entryID)

	return err

}
