package dashboard

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/dashboard"
	client "github.com/sea350/ustart_go/middleware/client"
	scrollpkg "github.com/sea350/ustart_go/properloading"
)

//AjaxLoadDashEntries ... pulls all entries for a given dashboard and fprints it back as json array (NOW WITH SCROLL!)
func AjaxLoadDashEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	wallID := r.FormValue("userID")
	dash, err := get.DashboardByUserID(client.Eclient, wallID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	res, entries, total, err := scrollpkg.ScrollPageDash(client.Eclient, dash.EntryIDs, "")
	if err != nil {
		fmt.Println(res)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	results := make(map[string]interface{})
	results["JournalEntries"] = entries
	results["ScrollID"] = res
	results["TotalHits"] = total

	data, err := json.Marshal(results)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	fmt.Fprintln(w, string(data))
}