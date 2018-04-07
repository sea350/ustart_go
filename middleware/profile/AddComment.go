package profile

import (
	"fmt"
	"net/http"

	uses "github.com/sea350/ustart_go/uses"
)

//AddComment ... Iunno
func AddComment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	postID := r.FormValue("followstat")
	postActual := postID[1:]
	comment := r.FormValue("commentz")
	id := r.FormValue("id") // userID
	fmt.Println("ID IS "+id)
	contentArray := []rune(comment)
	username := r.FormValue("username")
	fmt.Println(postActual+"is the post ID? ")
	err4 := uses.UserReplyEntry(eclient, id, postActual, contentArray)
	if err4 != nil {
		fmt.Println(err4)
	}
	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}

func AddComment2(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	postID := r.FormValue("postID")
	fmt.Println(postID)
	// postActual := postID[1:]
	comment := r.FormValue("body")
	contentArray := []rune(comment)
	fmt.Println(session.Values["DocID"].(string)+"IS PIKA")
	err4 := uses.UserReplyEntry(eclient, session.Values["DocID"].(string), postID, contentArray)
	if err4 != nil {
		fmt.Println(err4)
	}
	http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
}
