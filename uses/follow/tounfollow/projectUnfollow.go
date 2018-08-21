package uses

import (
	postFollow "github.com/sea350/ustart_go/post/follow"

	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectUnfollow ...
//allows for following a project, profile, or event
func ProjectUnfollow(eclient *elastic.Client, hostID string, viewerID string) error {
	//remove from following
	err := postFollow.RemoveProjectFollow(eclient, viewerID, "UserFollowing", hostID)
	if err != nil {
		return err
	}
	//remove from followers
	err = postFollow.RemoveProjectFollow(eclient, hostID, "Followers", viewerID)
	if err != nil {
		return err
	}

	return err
}
