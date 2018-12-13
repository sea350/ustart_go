package settings

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
			url, err := uses.UploadToS3(blob, test1.(string))
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
				return
			}
			err = uses.ChangeAccountImagesAndStatus(client.Eclient, session.Values["DocID"].(string), url, true, ``, "Avatar")
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			} else {
				session.Values["Avatar"] = url
				session.Save(r, w)
			}
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("invalid file upload")
		}
	case http.ErrMissingFile:
		blob := r.FormValue("image-data")
		//duplicate in AWS with docID as filename
		url, err := uses.UploadToS3(blob, test1.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
			return
		}
		err = uses.ChangeAccountImagesAndStatus(client.Eclient, session.Values["DocID"].(string), url, true, ``, "Avatar")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		} else {
			session.Values["Avatar"] = url
			session.Save(r, w)
		}
	default:
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	time.Sleep(2 * time.Second)
	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
}
