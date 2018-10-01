package settings

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/backend/get/event"
	client "github.com/sea350/ustart_go/backend/middleware/client"
	post "github.com/sea350/ustart_go/backend/post/event"
	uses "github.com/sea350/ustart_go/backend/uses"
)

//EventLogo ...
func EventLogo(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-image")
	blob := r.FormValue("image-data")

	//Getting eventID and member
	evntID := r.FormValue("eventID")
	evnt, member, err := get.EventAndMember(client.Eclient, r.FormValue("eventID"), test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	if uses.HasEventPrivilege("icon", evnt.PrivilegeProfiles, member) {
		//Checking if image is valid by checking the first 512 bytes for correct image signature
		buffer := make([]byte, 512)
		_, _ = clientFile.Read(buffer)
		defer clientFile.Close()
		if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
			err = post.UpdateEvent(client.Eclient, evntID, "Avatar", blob)
		}
	}

	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	return

}
