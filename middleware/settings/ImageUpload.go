package settings

import (
	"fmt"
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
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-image")

	blob := r.FormValue("image-data")
	if err != nil {
		fmt.Println("err: middleware/settings/imageupload line 23")
		fmt.Println(err)
		return
	}
	/*
		if header.Size == 0 {
			fmt.Println("warning: middleware/settings/imageupload file not sent")
		}
	*/
	fmt.Println("debug text: middleware/settings/imageupload line 33")
	fmt.Println(clientFile)

	//Checking if image is valid by checking the first 512 bytes for correct image signature

	buffer := make([]byte, 512)
	_, err = clientFile.Read(buffer)
	defer clientFile.Close()
	if err != nil {
		fmt.Println("warning: middleware/settings/imageupload line 42")
		fmt.Println(err)
	}
	fmt.Println(http.DetectContentType(buffer)[0:5])

	if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
		err = uses.ChangeAccountImagesAndStatus(eclient, session.Values["DocID"].(string), blob, true, ``, "Avatar")
		if err != nil {
			fmt.Println("err: middleware/settings/imageupload line 51")
			fmt.Println(err)
		}
	} else {
		fmt.Println("Warning: middleware/settings/imageupload invalid file upload")
	}

	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
}
