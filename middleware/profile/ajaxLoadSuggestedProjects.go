package profile

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/properloading"
)

//AjaxLoadSuggestedProjects ... Pulls suggested projects based on skills required
func AjaxLoadSuggestedProjects(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	uID := docID.(string)
	scrollID := r.FormValue("scrollID")

	var results = make(map[string]interface{})

	myUser, err := get.UserByID(client.Eclient, uID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	_, follDoc, err := getFollow.ByID(client.Eclient, uID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	sID, heads, hits, err := properloading.ScrollSuggestedProjects(client.Eclient, myUser.Tags, myUser.Projects, follDoc.ProjectFollowing, uID, scrollID)
	if err != nil && err != io.EOF {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	results["scrollID"] = sID
	results["SuggestedUsers"] = heads
	results["TotalHits"] = hits
	results["error"] = err

	data, err := json.Marshal(results)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
