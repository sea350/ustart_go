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
		return
	}
	r.ParseForm()
	blob := r.FormValue("banner-data")
	fmt.Println("debug text: middleware/settings/bannerupload line 21")
	fmt.Println(blob)

	err := post.UpdateUser(eclient, session.Values["DocID"].(string), "Banner", blob)
	if err != nil {
		fmt.Println("err: middleware/settings/bannerupload line 24")
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
	return

}
