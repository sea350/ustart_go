package settings

import (
	"net/http"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	uses "github.com/sea350/ustart_go/uses"
)

//BannerUpload ... pushes a new banner image into ES
func BannerUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()

	clientFile, header, err := r.FormFile("raw-banner")
	switch err {
	case nil:
		blob := r.FormValue("banner-data")
		buffer := make([]byte, 512)
		_, _ = clientFile.Read(buffer)
		defer clientFile.Close()

		usr, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		err = uses.DeleteFromS3(usr.Banner)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}

		url, err := uses.UploadToS3(blob, test1.(string)+"-"+time.Now().String()+"-banner")
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
			//Update the user banner
			err := post.UpdateUser(client.Eclient, session.Values["DocID"].(string), "Banner", url)
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}
		}
	case http.ErrMissingFile:
		blob := r.FormValue("banner-data")

		usr, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		err = uses.DeleteFromS3(usr.Banner)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}

		url, err := uses.UploadToS3(blob, test1.(string)+"-"+time.Now().String()+"-banner")
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		err = post.UpdateUser(client.Eclient, session.Values["DocID"].(string), "Banner", url)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
	default:

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	time.Sleep(2 * time.Second)
	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
}
