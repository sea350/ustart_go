package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	fmt.Fprintln(w, string(data))
}
