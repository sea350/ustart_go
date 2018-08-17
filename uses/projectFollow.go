package uses

import (
	postFollow "github.com/sea350/ustart_go/post/follow"

	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectFollow ...
//allows for following a project, profile, or event
func ProjectFollow(eclient *elastic.Client, hostID string, viewerID string) error {
	err := postFollow.NewProjectFollow(eclient, viewerID, "UserFollowing", hostID, false)
	if err != nil {
		return err
	}

	//false = append to followers
	err = postFollow.NewProjectFollow(eclient, hostID, "UserFollowers", viewerID, false)
	if err != nil {
		return err
	}

	return err
}
