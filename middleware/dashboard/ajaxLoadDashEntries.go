package dashboard

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	scrollpkg "github.com/sea350/ustart_go/properloading"
)

//AjaxLoadDashEntries ... pulls all entries for a given dashboard and fprints it back as json array (NOW WITH SCROLL!)
func AjaxLoadDashEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	wallID := r.FormValue("userID")
	/*
		dash, err := get.DashboardByUserID(client.Eclient, wallID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	*/

	//READ THIS:
	// _, followDoc, err := getFollow.ByID(client.Eclient, wallID)
	// if err != nil {
	// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 	log.Println(err)
	// }

	// followDoc.UserFollowing
	// followDoc.ProjectFollowing
	// followDoc.EventFollowing <-- might still be separate?
	//this is how to get the list of docIDs for the current pages following maps
	res, entries, total, err := scrollpkg.ScrollPageDash(client.Eclient, []string{wallID}, docID.(string), "")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	results := make(map[string]interface{})
	results["JournalEntries"] = entries
	results["ScrollID"] = res
	results["TotalHits"] = total

	data, err := json.Marshal(results)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
