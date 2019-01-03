package img

import (
	"io"
	
	"net/http"

	getImg "github.com/sea350/ustart_go/get/img"
	"github.com/sea350/ustart_go/middleware/client"
)

//Display ... draws only image
func Display(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[5:]

	img, err := getImg.ByID(client.Eclient, id)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	io.WriteString(w, img.Image)
}
