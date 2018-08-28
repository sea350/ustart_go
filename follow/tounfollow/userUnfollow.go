package uses

import (
	postFollow "github.com/sea350/ustart_go/post/follow"

	elastic "gopkg.in/olivere/elastic.v5"
)

//UserUnfollow ...
//allows for following a project, profile, or event
func UserUnfollow(eclient *elastic.Client, hostID string, viewerID string) error {
	//remove from following
	err := postFollow.RemoveUserFollow(eclient, viewerID, "UserFollowing", hostID, ``)
	if err != nil {
		return err
	}
	//remove from followers
	err = postFollow.RemoveUserFollow(eclient, hostID, "Followers", viewerID, ``)
	if err != nil {
		return err
	}

	return err
}
