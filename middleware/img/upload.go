package img

import (
	"encoding/base64"
	"os"

	"net/http"

	"strings"

	"github.com/sea350/ustart_go/middleware/client"
)

//Upload ... draws only image
func Upload(w http.ResponseWriter, r *http.Request) {
	data := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEX/TQBcNTh/AAAACklEQVR4nGNiAAAABgADNjd8qAAAAABJRU5ErkJggg=="
	// The actual image starts after the ","
	arr := []string{}
	i := strings.Index(data, ",")
	if i < 0 {
		client.Logger.Fatal("no comma")
	} else {
		arr = strings.Split(data, `,`)

	}
	// pass reader to NewDecoder
	//dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))
	dec, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		panic(err)
	}

	client.Logger.Println("debug text " + arr[1])
	//convert decoder to file
	f, err := os.Create("myfilename.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}

	// w.Header().Set("Content-Type", "image/png")
	// io.(w, dec)
}
