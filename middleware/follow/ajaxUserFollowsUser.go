package follow

import (
	
	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	getUser "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	postFollow "github.com/sea350/ustart_go/post/follow"
	post "github.com/sea350/ustart_go/post/notification"
	"github.com/sea350/ustart_go/types"
)

//AjaxUserFollowsUser ... an ajax call that toggles whether a user is actively following another user
func AjaxUserFollowsUser(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	ID, _ := session.Values["DocID"]
	if ID == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	followingID := r.FormValue("userID")
	if followingID == `` {
		return
	}

	_, err := getUser.UserByID(client.Eclient, followingID)
	if err != nil {
		return
	}

	isFollowing, err := getFollow.IsFollowing(client.Eclient, ID.(string), followingID, "user")
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	if !isFollowing {
		err = postFollow.NewUserFollow(client.Eclient, ID.(string), "following", followingID, false, "user")
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			return
		}

		err = postFollow.NewUserFollow(client.Eclient, followingID, "followers", ID.(string), false, "user")
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			return
		}

		var notif types.Notification
		notif.NewFollower(followingID, ID.(string))
		_, err := post.IndexNotification(client.Eclient, notif)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			return
		}
	} else {
		err = postFollow.RemoveUserFollow(client.Eclient, ID.(string), "following", followingID, "user")
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			return
		}

		err = postFollow.RemoveUserFollow(client.Eclient, followingID, "followers", ID.(string), "user")
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
	}

}
