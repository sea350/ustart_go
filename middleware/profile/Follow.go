package profile

import (
	"log"
	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Follow ... Iunno
func Follow(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		//No docID in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	userID := r.FormValue("userID")

	if test1.(string) == userID {
		return
	}

	isFollowed, err := getFollow.IsFollowing(client.Eclient, userID, session.Values["DocID"].(string), 2) //uses.IsFollowed(client.Eclient, userID, session.Values["DocID"].(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	if isFollowed == true {
		err := uses.UserUnfollow(client.Eclient, userID, session.Values["DocID"].(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	} else {
		err := uses.UserFollow(client.Eclient, userID, session.Values["DocID"].(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}

}
