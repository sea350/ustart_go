package settings

import (
	"fmt"
	"net/http"

	post "github.com/sea350/ustart_go/post/user"
)

//BannerUpload ... pushes a new banner image into ES
func BannerUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-banner")
	if err != nil {
		fmt.Println("err: middleware/settings/projectBannerUpload line 14\n", err)
	}

	blob := r.FormValue("banner-data")
	buffer := make([]byte, 512)
	_, _ = clientFile.Read(buffer)
	defer clientFile.Close()
	if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
		//Update the user banner
		err := post.UpdateUser(client.Eclient, session.Values["DocID"].(string), "Banner", blob)
		if err != nil {
			fmt.Println("err: middleware/settings/bannerUpload line 30\n", err)
		}
	}

	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
}
