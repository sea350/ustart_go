package follow

import (
	"fmt"
	"log"
	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	"github.com/sea350/ustart_go/middleware/client"
	postFollow "github.com/sea350/ustart_go/post/follow"
	post "github.com/sea350/ustart_go/post/notification"
	"github.com/sea350/ustart_go/types"
)

//AjaxUserFollowsProject ... an ajax call that toggles whether a user is actively following another user
func AjaxUserFollowsProject(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	ID, _ := session.Values["DocID"]
	if ID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	followingID := r.FormValue("projectID")
	if followingID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(`PROJECT ID NOT PASSED`)
		return
	}

	isFollowing, err := getFollow.IsFollowing(client.Eclient, ID.(string), followingID, "user")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	if !isFollowing {
		fmt.Println("TRYING TO FOLLOW")
		err = postFollow.NewUserFollow(client.Eclient, ID.(string), "following", followingID, false, "project")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}

		err = postFollow.NewProjectFollow(client.Eclient, followingID, "followers", ID.(string), false, "user")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}

		var notif types.Notification
		notif.NewFollower(followingID, ID.(string))
		_, err := post.IndexNotification(client.Eclient, notif)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}
	} else {
		fmt.Println("TRYING TO UNFOLLOW")
		err = postFollow.RemoveUserFollow(client.Eclient, ID.(string), "following", followingID, "user")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}

		err = postFollow.RemoveProjectFollow(client.Eclient, followingID, "followers", ID.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}

}
