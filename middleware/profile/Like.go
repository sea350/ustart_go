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
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	postID := r.FormValue("PostID")
	if postID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("WARNING: no post Id passed in")
		return
	}

	likeStatus, err := uses.IsLiked(client.Eclient, postID, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	if likeStatus == true {
		err := uses.UserUnlikeEntry(client.Eclient, postID, docID.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	} else {
		err := uses.UserLikeEntry(client.Eclient, postID, docID.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}
}
