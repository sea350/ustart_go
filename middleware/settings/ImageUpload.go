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
	blob := r.FormValue("image-data")
	// infile, header, clientFile := r.FormFile("raw-image")

	//Checking if image is valid by checking the first 512 bytes for correct image signature
	clientFile, _, _ := r.FormFile("raw-image")
	// defer clientFile.Close()
	buff := make([]byte, 512)
	if _, err = clientFile.Read(buff); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(http.DetectContentType(buff))

	err := uses.ChangeAccountImagesAndStatus(eclient, session.Values["DocID"].(string), blob, true, "hello", "Avatar")
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)
	return

}
