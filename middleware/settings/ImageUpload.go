package settings

import (
	"fmt"
	
	"net/http"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ImageUpload ...
func ImageUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/Settings/", http.StatusFound)
		return
	}
	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-image")
	switch err {
	case nil:
		//Checking if image is valid by checking the first 512 bytes for correct image signature
		blob := r.FormValue("image-data")
		buffer := make([]byte, 512)
		_, _ = clientFile.Read(buffer)
		defer clientFile.Close()
		if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
			//duplicate in AWS with docID as filename
			usr, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			}
			err = uses.DeleteFromS3(usr.Avatar)
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			}

			url, err := uses.UploadToS3(blob, test1.(string)+"-"+time.Now().String())
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
				http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
				return
			}
			err = uses.ChangeAccountImagesAndStatus(client.Eclient, session.Values["DocID"].(string), url, true, ``, "Avatar")
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			} else {
				session.Values["Avatar"] = url
				session.Save(r, w)
			}
		} else {
			
					client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"invalid file upload")
		}
	case http.ErrMissingFile:
		blob := r.FormValue("image-data")
		//duplicate in AWS with docID as filename
		usr, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}
		err = uses.DeleteFromS3(usr.Avatar)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}

		url, err := uses.UploadToS3(blob, test1.(string)+"-"+time.Now().String())
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
			http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
			return
		}
		err = uses.ChangeAccountImagesAndStatus(client.Eclient, session.Values["DocID"].(string), url, true, ``, "Avatar")
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		} else {
			session.Values["Avatar"] = url
			session.Save(r, w)
		}
	default:
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	time.Sleep(2 * time.Second)
	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
}
