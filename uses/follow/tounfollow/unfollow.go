package uses

// import (
// 	postFollow "github.com/sea350/ustart_go/post/follow"

// 	elastic "gopkg.in/olivere/elastic.v5"
// )

// //Unfollow ...
// //allows for following a project, profile, or event
// func Unfollow(eclient *elastic.Client, hostID string, viewerID string) error {
// 	//remove from following
// 	err := postFollow.RemoveFollow(eclient, viewerID, "Following", hostID)
// 	if err != nil {
// 		return err
// 	}
// 	//remove from followers
// 	err = postFollow.RemoveFollow(eclient, hostID, "Followers", viewerID)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }
