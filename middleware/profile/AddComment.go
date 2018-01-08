package profile 

import (
    "net/http"
    uses "github.com/sea350/ustart_go/uses"
    "fmt"
)


func AddComment(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
    if (test1 == nil){
    	// No username in session
    	http.Redirect(w, r, "/~", http.StatusFound)
    }

    r.ParseForm()
	postID := r.FormValue("followstat")
	postActual := postID[1:]
	comment := r.FormValue("commentz")
	id := r.FormValue("id") // userID
	contentArray := []rune(comment)
	username := r.FormValue("username")
	err4 := uses.UserReplyEntry(eclient,id,postActual,contentArray)
	if (err4 != nil){
		fmt.Println(err4)
	}
	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}



