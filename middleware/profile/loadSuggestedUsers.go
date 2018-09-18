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
	properloading "github.com/sea350/ustart_go/properloading"
	types "github.com/sea350/ustart_go/types"
)

//LoadSuggestedUsers ... pulls suggested users based on user's projects and shared skills
func LoadSuggestedUsers(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	ID := session.Values["DocID"].(string)
	scrollID := r.FormValue("scrollID")

	myUser, err := get.UserByID(client.Eclient, ID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	_, follDoc, err := getFollow.ByID(client.Eclient, ID)

	// var resArr []map[string]interface{}
	var resArr []types.FloatingHead
	count := 0
	for count < 3 {
		sID, heads, _, err := properloading.ScrollSuggestedUsers(client.Eclient, myUser.Tags, myUser.Projects, follDoc.UserFollowing, ID, scrollID)

		if err != nil && err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}

		if len(heads) > 0 {
			// var results = make(map[string]interface{})
			scrollID = sID
			// results["scrollID"] = sID
			// results["SuggestedUsers"] = heads

			resArr = append(resArr, heads...)
			count++
		}
		if err == io.EOF {
			break
		}
	}

	sendData := make(map[string]interface{})
	sendData["suggestions"] = resArr
	sendData["scrollID"] = scrollID

	data, err := json.Marshal(sendData)

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
