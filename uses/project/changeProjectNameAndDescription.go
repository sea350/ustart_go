package uses

import (
	post "github.com/sea350/ustart_go/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeProjectNameAndDescription ... CHANGES BOTH A PROJECT NAME AND DESCRIPTION
//Requires the target project's docID, new name and description
//Returns an error if there was a problem with database submission
//NOTE: it is possible Name change goes through but not Description
func ChangeProjectNameAndDescription(eclient *elastic.Client, projectID string, newName string, newDescription []rune) error {
	err := post.UpdateProject(eclient, projectID, "Name", newName)
	if err != nil {
		return err
	}
	err = post.UpdateProject(eclient, projectID, "Description", newDescription)
	return err
}
