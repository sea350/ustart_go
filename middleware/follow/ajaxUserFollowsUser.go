package profile

import (
	"log"
	"net/http"
	"os"

	getFollow "github.com/sea350/ustart_go/get/follow"
	"github.com/sea350/ustart_go/middleware/client"
	postFollow "github.com/sea350/ustart_go/post/follow"
)

//AjaxUserFollowsUser ... an ajax call that toggles whether a user is actively following another user
func AjaxUserFollowsUser(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	ID, _ := session.Values["DocID"]
	if ID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	var followingID string

	isFollowing, err := getFollow.IsFollowing(client.Eclient, ID.(string), followingID, "user")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return
	}

	if isFollowing {
		err = postFollow.NewUserFollow(client.Eclient, ID.(string), "following", followingID, false)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			return
		}

		err = postFollow.NewUserFollow(client.Eclient, followingID, "followers", ID.(string), false)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	} else {
		err = postFollow.RemoveUserFollow(client.Eclient, ID.(string), "following", followingID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			return
		}

		err = postFollow.RemoveUserFollow(client.Eclient, followingID, "followers", ID.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	}

}
