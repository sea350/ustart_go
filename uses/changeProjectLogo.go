package uses

import (
	post "github.com/sea350/ustart_go/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeProjectLogo ...
func ChangeProjectLogo(eclient *elastic.Client, projID string, newLogo string) error {

	return post.UpdateProject(eclient, projID, "Avatar", newLogo)

}
