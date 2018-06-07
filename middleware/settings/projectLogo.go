package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectLogo ...
func ProjectLogo(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	blob := r.FormValue("image-data")
	clientFile, header, err := r.FormFile("raw-image")

	//Getting project ID
	projID := r.FormValue("projectID")
	proj, err := get.ProjectByID(eclient, projID)
	if err != nil {
		fmt.Println(err)
	}

	buffer := make([]byte, 512)
	_, _ = clientFile.Read(buffer)
	defer clientFile.Close()
	if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
		err = uses.ChangeProjectLogo(eclient, projID, blob)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Warning: middleware/settings/projectLogo invalid file upload")
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return

}
