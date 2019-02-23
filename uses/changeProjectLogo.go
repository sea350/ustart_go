package uses

import (
	post "github.com/sea350/ustart_go/post/project"
	elastic "github.com/olivere/elastic"
)

//ChangeProjectLogo ...
func ChangeProjectLogo(eclient *elastic.Client, projID string, newLogo string) error {

	return post.UpdateProject(eclient, projID, "Avatar", newLogo)

}
