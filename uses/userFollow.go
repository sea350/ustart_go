package uses

import (
	"time"

	get "github.com/sea350/ustart_go/get/notification"
	postFollow "github.com/sea350/ustart_go/post/follow"
	post "github.com/sea350/ustart_go/post/notification"
	types "github.com/sea350/ustart_go/types"
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

	newNotif := types.Notification{Class: 4, DocID: hostID, Timestamp: time.Now(), ReferenceIDs: []string{viewerID}}

	notifID, err := post.IndexNotification(eclient, newNotif)
	if err != nil {
		return err
	}

	proxyID, err := get.ProxyIDByUserID(eclient, viewerID)
	if err != nil {
		return err
	}

	err = post.AppendToProxy(eclient, proxyID, notifID, true)

	return err
}
