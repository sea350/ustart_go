package uses

import (
	"time"

	get "github.com/sea350/ustart_go/get/notification"
	post "github.com/sea350/ustart_go/post/notification"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//FollowUserNotif ...
func FollowUserNotif(eclient *elastic.Client, followedUserID string, followerID string) error {
	newNotif := types.Notification{Class: 4, DocID: followedUserID, Timestamp: time.Now(), ReferenceIDs: []string{followerID}}

	notifID, err := post.IndexNotification(eclient, newNotif)
	if err != nil {
		return err
	}

	proxyID, err := get.ProxyIDByUserID(eclient, followedUserID)
	if err != nil {
		return err
	}

	err = post.AppendToProxyNotification(eclient, proxyID, notifID)

	return err
}
