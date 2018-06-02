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
	clientFile, _, err := r.FormFile("raw-image")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("debug text: middleware/settings/imageupload line 26")
	fmt.Println(clientFile)
	blob := r.FormValue("image-data")

	//Checking if image is valid by checking the first 512 bytes for correct image signature
	defer clientFile.Close()
	buffer := make([]byte, 512)
	_, err = clientFile.Read(buffer)
	if err != nil {
		fmt.Println("err: middleware/settings/imageupload line 33")
		fmt.Println(err)
		return
	}
	//fmt.Println(http.DetectContentType(buffer)[0:5])

	if http.DetectContentType(buffer)[0:5] == "image" {
		err = uses.ChangeAccountImagesAndStatus(eclient, session.Values["DocID"].(string), blob, true, "hello", "Avatar")
		if err != nil {
			fmt.Println("err: middleware/settings/imageupload line 42")
			fmt.Println(err)
		}
	}

	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
	return

}
