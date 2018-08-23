package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/event"
	post "github.com/sea350/ustart_go/post/event"
	uses "github.com/sea350/ustart_go/uses"
)

//EventBannerUpload ... pushes a new banner image into ES
func EventBannerUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	//maybe uncomment later:
	// if test1 == nil {
	// 	fmt.Println(test1)
	// 	http.Redirect(w, r, "/~", http.StatusFound)
	// 	return
	// }

	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-banner")
	blob := r.FormValue("banner-data")

	//Get event by ID
	// evnt, err := get.EventbyID(eclient, r.FormValue("eventID"))
	// if err != nil {
	// 	fmt.Println("err: middleware/settings/eventBannerUpload line 33\n", err)
	// }

	//get the member
	evnt, member, err := get.EventAndMember(eclient, r.FormValue("eventID"), test1.(string))
	//check privilege
	if uses.HasEventPrivilege("banner", evnt.PrivilegeProfiles, member) {
		buffer := make([]byte, 512)
		_, _ = clientFile.Read(buffer)
		defer clientFile.Close()
		if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
			//Update the event banner
			err = post.UpdateEvent(eclient, r.FormValue("eventID"), "Banner", blob)
			if err != nil {
				fmt.Println("err: middleware/settings/eventbannerupload line 50\n", err)
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
