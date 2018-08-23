package uses

import (
	postFollow "github.com/sea350/ustart_go/post/follow"

	elastic "gopkg.in/olivere/elastic.v5"
)

//Follow ...
//allows for following a project, profile, or event
func Follow(eclient *elastic.Client, hostID string, viewerID string) error {
	err := postFollow.NewUserFollow(eclient, viewerID, "Following", hostID, false)
	if err != nil {
		return err
	}

	//false = append to followers
	err = postFollow.NewUserFollow(eclient, hostID, "Followers", viewerID, false)
	if err != nil {
		return err
	}

	return err
}
