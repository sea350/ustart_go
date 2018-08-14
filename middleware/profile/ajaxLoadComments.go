package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//AjaxLoadComments ... pulls all entries for a given user and fprints it back as a json array
func AjaxLoadComments(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	parentID := r.FormValue("postID")
	_, entries, err := uses.LoadComments(client.Eclient, parentID, docID.(string), 0, -1)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
