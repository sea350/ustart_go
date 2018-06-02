package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	post "github.com/sea350/ustart_go/post/project"
)

//ProjectBannerUpload ... pushes a new banner image into ES
func ProjectBannerUpload(w http.ResponseWriter, r *http.Request) {
	//session, _ := store.Get(r, "session_please")
	/*test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}*/
	r.ParseForm()
	blob := r.FormValue("banner-data")

	proj, err := get.ProjectByID(eclient, r.FormValue("projectID"))
	if err != nil {
		fmt.Println(err)
	}

	err = post.UpdateProject(eclient, r.FormValue("projectID"), "Banner", blob)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return

}
