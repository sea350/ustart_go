package profile

import (
	"log"
	"net/http"
	"os"

	getFollow "github.com/sea350/ustart_go/get/follow"
	"github.com/sea350/ustart_go/middleware/client"
	postFollow "github.com/sea350/ustart_go/post/follow"
)

//AjaxUserFollowsUser ... an ajax call that changes whether a project is visible on the user page
func AjaxUserFollowsUser(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	ID, _ := session.Values["DocID"]
	if ID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	eventID := r.FormValue("eventID")
	var followingID string

	isFollowing, err := getFollow.IsFollowing(client.Eclient, ID, followingID, "user")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return
	}

	isFollowedBy, err := getFollow.IsFollowedBy(client.Eclient, followingID, ID, "user")

	if isFollowing {
		err = postFollow.NewUserFollow(client.Eclient, ID, "following", followingID, false)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			return
		}

		err = postFollow.NewUserFollow(client.Eclient, followingID, "followers", ID, false)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	}

}
