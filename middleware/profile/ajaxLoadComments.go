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
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	parentID := r.FormValue("postID")
	_, entries, err := uses.LoadComments(client.Eclient, parentID, 0, -1)
	if err != nil {
		log.Println("Error: middleware/profile/ajaxLoadComments line 23")
		log.Println(err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		log.Println("Error: middleware/profile/ajaxloadcomments line 30")
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
