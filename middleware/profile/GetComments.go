package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//GetComments ... gets comments???
func GetComments(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		//No docid in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()
	postID := r.FormValue("PostID")
	_, arrayofComments, err := uses.LoadComments(client.Eclient, postID, docID.(string), 0, -1)
	if err != nil && err != errors.New("This entry is not visible") {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	data, err := json.Marshal(arrayofComments)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(data))
}
