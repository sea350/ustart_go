package uses

import (
	post "github.com/sea350/ustart_go/post/project"
	elastic "github.com/olivere/elastic"
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
