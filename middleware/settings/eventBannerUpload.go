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
	//maybe uncomment later:
	// if test1 == nil {
	// 	fmt.Println(test1)
	// 	http.Redirect(w, r, "/~", http.StatusFound)
	// 	return
	// }

	r.ParseForm()

	//get the member
	evnt, member, err := get.EventAndMember(client.Eclient, r.FormValue("eventID"), test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
		return
	}

	//Get image upload
	clientFile, header, err := r.FormFile("raw-banner")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
		return
	}
	blob := r.FormValue("banner-data")

	//check privilege
	if uses.HasEventPrivilege("banner", evnt.PrivilegeProfiles, member) {
		buffer := make([]byte, 512)
		_, _ = clientFile.Read(buffer)
		defer clientFile.Close()
		if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
			//Update the event banner
			err = post.UpdateEvent(client.Eclient, r.FormValue("eventID"), "Banner", blob)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		} else {
			fmt.Println("err: middleware/settings/eventbannerupload invalid file upload")
		}

	} else {
		fmt.Println("err: middleware/settings/eventLogo  you have no permission to change event banner")
	}
	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	return
}
