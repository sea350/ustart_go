package uses

import (
	post "github.com/sea350/ustart_go/backend/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectNameAndDescription ...
func ProjectNameAndDescription(eclient *elastic.Client, projID string, name string, desc []rune) error {

	err := post.UpdateProject(eclient, projID, "Name", name)
	if err != nil {
		return err
	}
	err = post.UpdateProject(eclient, projID, "Description", desc)
	return err

}
