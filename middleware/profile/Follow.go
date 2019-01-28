package profile

import (
	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	client "github.com/sea350/ustart_go/middleware/client"
)

//Follow ... Iunno
func Follow(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		//No docID in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	userID := r.FormValue("userID")

	if test1.(string) == userID {
		return
	}
	//check if following
	isFollowed, err := getFollow.IsFollowing(client.Eclient, userID, session.Values["DocID"].(string), "user") //uses.IsFollowed(client.Eclient, userID, session.Values["DocID"].(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	if isFollowed == true {
		// err := usesFollow.UserUnfollow(client.Eclient, userID, session.Values["DocID"].(string))
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
	} else {
		// err := uses.UserFollow(client.Eclient, userID, session.Values["DocID"].(string))
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
	}

}
