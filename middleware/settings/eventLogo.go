package settings

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
	uses "github.com/sea350/ustart_go/uses"
)

//EventLogo ... Uploads logo for an event, checks for propper permissions
func EventLogo(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()

	//Get uploaded image
	clientFile, header, err := r.FormFile("raw-image")

	//get the eventID and member, have to put this after getting image upload or it wont work (idk why)
	evnt, member, err1 := get.EventAndMember(client.Eclient, r.FormValue("eventID"), test1.(string))
	if err1 != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	blob := r.FormValue("image-data")

	switch err {
	case nil:
		if uses.HasEventPrivilege("icon", evnt.PrivilegeProfiles, member) {
			//Checking if image is valid by checking the first 512 bytes for correct image signature
			buffer := make([]byte, 512)
			_, _ = clientFile.Read(buffer)
			defer clientFile.Close()
			if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
				err = post.UpdateEvent(client.Eclient, r.FormValue("eventID"), "Avatar", blob)
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					log.Println(err)
				}
			}
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("You do not have permission to change event logo")
		}
	case http.ErrMissingFile:
		//If file is not uploading (resizing original image)
		if uses.HasEventPrivilege("icon", evnt.PrivilegeProfiles, member) {
			err = post.UpdateEvent(client.Eclient, r.FormValue("eventID"), "Avatar", blob)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}

		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("You do not have permission to change event logo")
		}
	default:
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
}
