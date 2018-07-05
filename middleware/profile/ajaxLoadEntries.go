package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//AjaxLoadEntries ... pulls all entries for a given array of entry ids and fprints it back as a json array
func AjaxLoadEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	entryIDsString := r.FormValue("postIndex")
	entryIDsArr := uses.ConvertStrToStrArr(entryIDsString)

	entries, err := uses.LoadEntries(client.Eclient, entryIDsArr)
	if err != nil {
		log.Println("Error: middleware/profile/ajaxLoadEntries line 25")
		log.Println(err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		log.Println("Error: middleware/profile/ajaxloadentries line 31")
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
