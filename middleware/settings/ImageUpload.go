package settings

import (
	"fmt"
	"log"
	"net/http"

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
	blob := r.FormValue("image-data")
	clientFile, header, err := r.FormFile("raw-image")
	switch err {
	case nil:
		fmt.Println("------------------------CASE 1------------------------")
		//Checking if image is valid by checking the first 512 bytes for correct image signature
		buffer := make([]byte, 512)
		_, _ = clientFile.Read(buffer)
		defer clientFile.Close()
		if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
			fmt.Println("------------------------CASE 1A------------------------")
			fmt.Println("Blob== ", blob)
			err = uses.ChangeAccountImagesAndStatus(client.Eclient, session.Values["DocID"].(string), blob, true, ``, "Avatar")
			if err != nil {
				fmt.Println("------------------------CASE 1AA------------------------")
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			} else {
				fmt.Println("------------------------CASE 1AB------------------------")
				session.Values["Avatar"] = blob
				session.Save(r, w)
			}
		} else {
			fmt.Println("------------------------CASE 1B------------------------")
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("invalid file upload")
		}
	case http.ErrMissingFile:
		fmt.Println("------------------------CASE 2------------------------")

		err = uses.ChangeAccountImagesAndStatus(client.Eclient, session.Values["DocID"].(string), blob, true, ``, "Avatar")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		} else {
			session.Values["Avatar"] = blob
			session.Save(r, w)
		}
		http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
	default:
		fmt.Println("------------------------CASE 3------------------------")

		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
	}

	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
}
