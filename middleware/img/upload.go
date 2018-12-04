package img

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"
)

//Upload ... draws only image
func Upload(w http.ResponseWriter, r *http.Request) {
	data := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEX/TQBcNTh/AAAACklEQVR4nGNiAAAABgADNjd8qAAAAABJRU5ErkJggg=="
	// The actual image starts after the ","
	i := strings.Index(data, ",")
	if i < 0 {
		log.Fatal("no comma")
	}
	// pass reader to NewDecoder
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))
	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, dec)
}
