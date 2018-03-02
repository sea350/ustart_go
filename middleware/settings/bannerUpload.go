package settings

import (
	"fmt"
	"net/http"

	post "github.com/sea350/ustart_go/post/user"
)

//BannerUpload ... pushes a new banner image into ES
func BannerUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	r.ParseForm()
	blob := r.FormValue("banner-data")

	err := post.UpdateUser(eclient, session.Values["DocID"].(string), "Banner", blob)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}
