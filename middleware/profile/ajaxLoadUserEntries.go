package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
)

//AjaxLoadUserEntries ... pulls all entries for a given user and fprints it back as a json array
func AjaxLoadUserEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	wallID := r.FormValue("userID")
	user, err := get.UserByID(client.Eclient, wallID)
	if err != nil {
		log.Println("Error: middleware/profile/ajaxLoadUserEntries line 24")
		log.Println(err)
	}
	entries, err := uses.LoadEntries(client.Eclient, user.EntryIDs)
	if err != nil {
		log.Println("Error: middleware/profile/ajaxLoadUserEntries line 29")
		log.Println(err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		log.Println("Error: middleware/profile/ajaxLoadUserEntries line 35")
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
