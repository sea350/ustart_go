package profile

import (
	"fmt"
	"net/http"

	uses "github.com/sea350/ustart_go/uses"
)

//Follow ... Iunno
func Follow(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		//No docID in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	fname := r.FormValue("userID")
	fmt.Println("debug text middleware/profile/follow, this is fname:")
	fmt.Println(fname)
	fmt.Println("debug text middleware/profile/follow, this is docID:")
	fmt.Println(session.Values["DocID"].(string))

	isFollowed, err4 := uses.IsFollowed(eclient, session.Values["DocID"].(string), fname)
	if err4 != nil {
		fmt.Println("this is an error (Follow.go: 24)")
		fmt.Println(err4)
	}
	if isFollowed == true {
		err := uses.UserUnfollow(eclient, session.Values["DocID"].(string), fname)
		if err != nil {
			fmt.Println("this is an error (Follow.go: 30)")
			fmt.Println(err)
		}
	} else {
		err := uses.UserFollow(eclient, session.Values["DocID"].(string), fname)
		if err != nil {
			fmt.Println("this is an error (Follow.go: 36)")
			fmt.Println(err)
		}
	}

}
