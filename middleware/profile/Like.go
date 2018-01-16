package profile

import (
	"fmt"
	"net/http"

	uses "github.com/sea350/ustart_go/uses"
)

//Like ... Iunno
func Like(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	r.ParseForm()
	postid := r.FormValue("PostID")
	postactual := postid[10:]       // postid has to be trimmed
	docid := r.FormValue("selfDoc") // docid of the doc you are viewing double check
	likeStatus, err4 := uses.IsLiked(eclient, postactual, docid)
	if err4 != nil {
		fmt.Println(err4)
	}
	if likeStatus == true {
		err := uses.UserUnlikeEntry(eclient, postactual, docid)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := uses.UserLikeEntry(eclient, postactual, docid)
		if err != nil {
			fmt.Println(err)
		}
	}
}
