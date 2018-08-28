package follow

import (
	"fmt"
	"log"
	"net/http"
	"os"

	getFollow "github.com/sea350/ustart_go/get/follow"
	"github.com/sea350/ustart_go/middleware/client"
	postFollow "github.com/sea350/ustart_go/post/follow"
)

//AjaxUserFollowsProject ... an ajax call that changes whether a user is actively following a project
func AjaxUserFollowsProject(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	ID, _ := session.Values["DocID"]
	if ID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	followingID := r.FormValue("projectID")

	isFollowing, err := getFollow.IsFollowing(client.Eclient, ID.(string), followingID, "project")
	fmt.Println("AJAX isFollowing:", isFollowing)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return
	}

	if !isFollowing {

		err = postFollow.NewUserFollow(client.Eclient, ID.(string), "following", followingID, false, "project")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			return
		}

		err = postFollow.NewProjectFollow(client.Eclient, followingID, "followers", ID.(string), false, "user")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	} else {

		err = postFollow.RemoveUserFollow(client.Eclient, ID.(string), "following", followingID, "project")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			return
		}

		err = postFollow.RemoveProjectFollow(client.Eclient, followingID, "followers", ID.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	}

}
