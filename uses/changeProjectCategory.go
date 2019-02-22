package uses

import (
	post "github.com/sea350/ustart_go/post/project"
	elastic "github.com/olivere/elastic"
)

//ChangeProjectCategory ... CHANGES PROJECT THE PROJECT'S CATEGORY
//Requires the target project's docID, all aspects of a types.LocStruct
//Returns an error if there was a problem with database submission
func ChangeProjectCategory(eclient *elastic.Client, projectID string, category string) error {
	err := post.UpdateProject(eclient, projectID, "Category", category)
	return err
}
