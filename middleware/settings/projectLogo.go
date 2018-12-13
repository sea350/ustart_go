package settings

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectLogo Upload a new Project image icon
func ProjectLogo(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-image")

	proj, member, err1 := get.ProjAndMember(client.Eclient, r.FormValue("projectID"), test1.(string))
	if err1 != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err, "Project or Member not found")
		http.Redirect(w, r, "/ProjectSettings/"+proj.URLName, http.StatusFound)

	}

	switch err {
	case nil:
		blob := r.FormValue("image-data")
		if uses.HasPrivilege("icon", proj.PrivilegeProfiles, member) {
			//Checking if image is valid by checking the first 512 bytes for correct image signature
			buffer := make([]byte, 512)
			_, _ = clientFile.Read(buffer)
			defer clientFile.Close()
			if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
				url, err := uses.UploadToS3(blob, r.FormValue("projectID"))
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					log.Println(err)
					http.Redirect(w, r, "/ProjectSettings/"+proj.URLName, http.StatusFound)
					return
				}
				err = uses.ChangeProjectLogo(client.Eclient, r.FormValue("projectID"), url)
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					log.Println(err)
				}
			} else {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err, "You do not have permission to change project logo")
		}
	case http.ErrMissingFile:
		blob := r.FormValue("image-data")
		if uses.HasPrivilege("icon", proj.PrivilegeProfiles, member) {
			url, err := uses.UploadToS3(blob, r.FormValue("projectID"))
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				http.Redirect(w, r, "/ProjectSettings/"+proj.URLName, http.StatusFound)
				return
			}
			err = uses.ChangeProjectLogo(client.Eclient, r.FormValue("projectID"), url)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err, "You do not have permission to change project logo")
		}
	default:
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	http.Redirect(w, r, "/ProjectSettings/"+proj.URLName, http.StatusFound)
}
