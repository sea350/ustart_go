package settings

import (
	"net/http"
	"time"

	client "github.com/sea350/ustart_go/middleware/client"

	get "github.com/sea350/ustart_go/get/event"
	post "github.com/sea350/ustart_go/post/event"
	uses "github.com/sea350/ustart_go/uses"
)

//EventBannerUpload ... pushes a new banner image into ES
func EventBannerUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()

	//Get image upload
	clientFile, header, err := r.FormFile("raw-banner")
	//get the eventID and member, have to put this after getting image upload or it wont work (idk why)
	evnt, member, err1 := get.EventAndMember(client.Eclient, r.FormValue("eventID"), test1.(string))
	if err1 != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err1)
		http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	}

	switch err {
	case nil:
		blob := r.FormValue("banner-data")
		//check privilege
		if uses.HasEventPrivilege("banner", evnt.PrivilegeProfiles, member) {
			buffer := make([]byte, 512)
			_, _ = clientFile.Read(buffer)
			defer clientFile.Close()
			if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {

				err = uses.DeleteFromS3(evnt.Banner)
				if err != nil {

					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
				}

				url, err := uses.UploadToS3(blob, r.FormValue("eventID")+"-"+time.Now().String()+"-banner")
				if err != nil {

					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
					http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
					return
				}
				err = post.UpdateEvent(client.Eclient, r.FormValue("eventID"), "Banner", url)
				if err != nil {

					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
				}
			} else {

				client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Invalid file upload")
			}
		} else {

			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | err: " + "Unknown error encountered")
		}
	case http.ErrMissingFile:
		//If file is not uploaded
		blob := r.FormValue("banner-data")
		if uses.HasEventPrivilege("banner", evnt.PrivilegeProfiles, member) {
			err = uses.DeleteFromS3(evnt.Banner)
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			}

			url, err := uses.UploadToS3(blob, r.FormValue("eventID")+"-"+time.Now().String()+"-banner")
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
				http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
				return
			}
			err = post.UpdateEvent(client.Eclient, r.FormValue("eventID"), "Banner", url)
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			}
		} else {

			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "You do not have permission to change event banner")
		}
	default:

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	time.Sleep(2 * time.Second)
	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
}
