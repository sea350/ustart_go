package settings

import (
	"fmt"
	"log"
	"net/http"

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
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()

	//Get image upload
	clientFile, header, err := r.FormFile("raw-banner")
	//get the member and member, have to put this after getting image upload or it wont work (idk why)
	evnt, member, err1 := get.EventAndMember(client.Eclient, r.FormValue("eventID"), test1.(string))
	if err1 != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err1)
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
				err = post.UpdateEvent(client.Eclient, r.FormValue("eventID"), "Banner", blob)
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					log.Println(err)
				}
			} else {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println("Invalid file upload")
			}
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println()
		}
	case http.ErrMissingFile:
		//If file is not uploaded
		blob := r.FormValue("banner-data")
		if uses.HasEventPrivilege("banner", evnt.PrivilegeProfiles, member) {
			err = post.UpdateEvent(client.Eclient, r.FormValue("eventID"), "Banner", blob)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("You do not have permission to change event banner")
		}
	default:
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
}
