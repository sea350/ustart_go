package settings

import (
	
	"net/http"
	"time"

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
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
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
				err = uses.DeleteFromS3(evnt.Avatar)
				if err != nil {
					
					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
				}

				url, err := uses.UploadToS3(blob, r.FormValue("eventID")+"-"+time.Now().String())
				if err != nil {
					
					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
					http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
					return
				}
				err = post.UpdateEvent(client.Eclient, r.FormValue("eventID"), "Avatar", url)
				if err != nil {
					
					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
				}
			}
		} else {
			
					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"You do not have permission to change event logo")
		}
	case http.ErrMissingFile:
		//If file is not uploading (resizing original image)
		if uses.HasEventPrivilege("icon", evnt.PrivilegeProfiles, member) {
			err = uses.DeleteFromS3(evnt.Avatar)
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}

			url, err := uses.UploadToS3(blob, r.FormValue("eventID")+"-"+time.Now().String())
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
				http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
				return
			}
			err = post.UpdateEvent(client.Eclient, r.FormValue("eventID"), "Avatar", url)
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}

		} else {
			
					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"You do not have permission to change event logo")
		}
	default:
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	time.Sleep(2 * time.Second)
	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
}
