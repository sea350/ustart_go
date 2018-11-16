package dashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	client "github.com/sea350/ustart_go/middleware/client"
	scrollpkg "github.com/sea350/ustart_go/properloading"
)

//AjaxLoadDashEntries ... pulls all entries for a given dashboard and fprints it back as json array (NOW WITH SCROLL!)
func AjaxLoadDashEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	// wallID := r.FormValue("userID")

	//READ THIS:
	_, followDoc, err := getFollow.ByID(client.Eclient, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	var docs []string
	for i := range followDoc.UserFollowing {
		docs = append(docs, i)

	}
	for i := range followDoc.ProjectFollowing {
		docs = append(docs, i)

	}
	sID := r.FormValue("scrollID")
	// followDoc.UserFollowing
	// followDoc.ProjectFollowing
	// followDoc.EventFollowing <-- might still be separate?
	//this is how to get the list of docIDs for the current pages following maps
	res, entries, total, err := scrollpkg.ScrollPageDash(client.Eclient, docs, docID.(string), sID)
	if err != nil && err != io.EOF {
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

	// cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ScrollID: res, Wall: entries, Hits: total}

	fmt.Fprintln(w, string(data))
}
