package settings

import (
	"fmt"
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectBannerUpload ... pushes a new banner image into ES
func ProjectBannerUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]

	r.ParseForm()

	//get the member
	proj, member, err := get.ProjAndMember(client.Eclient, r.FormValue("projectID"), test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err, "Project or Member not found")
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	blob := r.FormValue("banner-data")
	clientFile, header, err := r.FormFile("raw-banner")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	//check privilege
	if uses.HasPrivilege("banner", proj.PrivilegeProfiles, member) {
		buffer := make([]byte, 512)
		_, _ = clientFile.Read(buffer)
		defer clientFile.Close()
		if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
			//Update the project banner
			err = post.UpdateProject(client.Eclient, r.FormValue("projectID"), "Banner", blob)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		} else {
			fmt.Println("err: middleware/settings/projectBannerUpload invalid file upload")
		}

	} else {

		fmt.Println("err: middleware/settings/projectLogo  you have no permission to change project banner")
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return
}
