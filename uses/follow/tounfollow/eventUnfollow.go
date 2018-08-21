package uses

import (
	postFollow "github.com/sea350/ustart_go/post/follow"

	elastic "gopkg.in/olivere/elastic.v5"
)

//EventUnfollow ...
//allows for following a project, profile, or event
func EventUnfollow(eclient *elastic.Client, hostID string, viewerID string) error {
	//remove from following
	err := postFollow.RemoveEventFollow(eclient, viewerID, "EventFollowing", hostID)
	if err != nil {
		return err
	}
	//remove from followers
	err = postFollow.RemoveEventFollow(eclient, hostID, "EventFollowers", viewerID)
	if err != nil {
		return err
	}

	return err
}
