package uses

import (
	post "github.com/sea350/ustart_go/backend/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeProjectCategory ... CHANGES PROJECT THE PROJECT'S CATEGORY
//Requires the target project's docID, all aspects of a types.LocStruct
//Returns an error if there was a problem with database submission
func ChangeProjectCategory(eclient *elastic.Client, projectID string, category string) error {
	err := post.UpdateProject(eclient, projectID, "Category", category)
	return err
}
