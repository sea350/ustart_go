package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	post "github.com/sea350/ustart_go/post/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectBannerUpload ... pushes a new banner image into ES
func ProjectBannerUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	//maybe uncomment later:
	// if test1 == nil {
	// 	fmt.Println(test1)
	// 	http.Redirect(w, r, "/~", http.StatusFound)
	// 	return
	// }

	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-banner")
	if err != nil {
		fmt.Println("err: middleware/settings/projectBannerUpload line 26\n", err)
	}
	blob := r.FormValue("banner-data")

	//Get project by ID
	// proj, err := get.ProjectByID(client.Eclient, r.FormValue("projectID"))
	// if err != nil {
	// 	fmt.Println("err: middleware/settings/projectbannerupload line 33\n", err)
	// }

	//get the member
	proj, member, err := get.ProjAndMember(client.Eclient, r.FormValue("projectID"), test1.(string))
	if err != nil {
		fmt.Println("err: middleware/settings/projectbannerupload line 40\n", err)
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
				fmt.Println("err: middleware/settings/projectbannerupload line 50\n", err)
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
