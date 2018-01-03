package profile 

import (
    "net/http"
    uses "github.com/sea350/ustart_go/uses"
    "fmt"
)

func Follow(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
    	//No docID in session  
    http.Redirect(w, r, "/~", http.StatusFound)
    }

	r.ParseForm()
	fname := r.FormValue("userID")
	//following := r.FormValue("Following")

	isFollowed, err4 := uses.IsFollowed(eclient, session.Values["DocID"].(string),fname)
	if (err4 != nil){
		fmt.Println(err4)
	}
	if (isFollowed == true){
		err := uses.UserUnfollow(eclient,session.Values["DocID"].(string),fname)
		if (err != nil){
			fmt.Println(err);
		}
	}else{
		err := uses.UserFollow(eclient,session.Values["DocID"].(string),fname)
		if (err != nil){
			fmt.Println(err);
		}	
	}

}
