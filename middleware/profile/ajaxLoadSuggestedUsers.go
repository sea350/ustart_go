package profile

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	getFollow "github.com/sea350/ustart_go/get/follow"
	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	properloading "github.com/sea350/ustart_go/properloading"
)

//AjaxLoadSuggestedUsers ... pulls suggested users based on user's projects and shared skills
func AjaxLoadSuggestedUsers(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	uID := docID.(string)
	scrollID := r.FormValue("scrollID")

	var results = make(map[string]interface{})

	myUser, err := get.UserByID(client.Eclient, uID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		results["error"] = err
		data, err := json.Marshal(results)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}

		fmt.Fprintln(w, string(data))
		return
	}

	_, follDoc, err := getFollow.ByID(client.Eclient, uID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		results["error"] = err
		data, err := json.Marshal(results)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}

		fmt.Fprintln(w, string(data))
		return
	}

	sID, heads, hits, err := properloading.ScrollSuggestedUsers(client.Eclient, myUser.Class, myUser.Tags, myUser.Projects, follDoc.UserFollowing, uID, scrollID, myUser.Majors, myUser.University)
	if err != nil && err != io.EOF {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	results["scrollID"] = sID
	results["SuggestedUsers"] = heads
	results["TotalHits"] = hits
	results["error"] = err

	data, err := json.Marshal(results)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	fmt.Fprintln(w, string(data))
}
