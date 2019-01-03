package settings

import (
	"net/http"
	"time"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectBannerUpload ... pushes a new banner image into ES
func ProjectBannerUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-banner")

	//Get the project and member
	proj, member, err1 := get.ProjAndMember(client.Eclient, r.FormValue("projectID"), test1.(string))
	if err1 != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s Project or Member not found", err)
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	switch err {
	case nil:
		blob := r.FormValue("banner-data")
		if uses.HasPrivilege("banner", proj.PrivilegeProfiles, member) {
			buffer := make([]byte, 512)
			_, _ = clientFile.Read(buffer)
			defer clientFile.Close()
			if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
				//Update the project banner
				err = uses.DeleteFromS3(proj.Banner)
				if err != nil {

					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
				}

				url, err := uses.UploadToS3(blob, r.FormValue("projectID")+"-"+time.Now().String()+"-banner")
				if err != nil {

					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
				}
				err = post.UpdateProject(client.Eclient, r.FormValue("projectID"), "Banner", url)
				if err != nil {

					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
				}
			} else {

				client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Invalid file upload")
			}
		} else {

			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "You do not have permission to change event banner")
		}
	case http.ErrMissingFile:
		blob := r.FormValue("banner-data")
		if uses.HasPrivilege("banner", proj.PrivilegeProfiles, member) {
			//Update the project banner
			err = uses.DeleteFromS3(proj.Banner)
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			}

			url, err := uses.UploadToS3(blob, r.FormValue("projectID")+"-"+time.Now().String()+"-banner")
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			}
			err = post.UpdateProject(client.Eclient, r.FormValue("projectID"), "Banner", url)
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			} else {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s Invalid file upload", err)
			}
		} else {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s You do not have permission to change event banner", err)
		}

	default:

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	time.Sleep(2 * time.Second)
	http.Redirect(w, r, "/ProjectSettings/"+proj.URLName, http.StatusFound)
}
