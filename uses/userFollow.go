package uses

import (
	postFollow "github.com/sea350/ustart_go/post/follow"

	elastic "gopkg.in/olivere/elastic.v5"
)

//UserFollow ...
//allows for following a project, profile, or event
func UserFollow(eclient *elastic.Client, hostID string, viewerID string) error {
	err := postFollow.NewUserFollow(eclient, viewerID, "UserFollowing", hostID, false)
	if err != nil {
		return err
	}

	//false = append to followers
	err = postFollow.NewUserFollow(eclient, hostID, "UserFollowers", viewerID, false)
	if err != nil {
		return err
	}

	return err
}
