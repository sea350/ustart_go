package profile

import (
	"log"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Like ... likes a post, designed for ajax
func Like(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	postid := r.FormValue("PostID")

	likeStatus, err := uses.IsLiked(client.Eclient, postid, session.Values["DocID"].(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	if likeStatus == true {
		err := uses.UserUnlikeEntry(client.Eclient, postid, session.Values["DocID"].(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	} else {
		err := uses.UserLikeEntry(client.Eclient, postid, session.Values["DocID"].(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}
}
