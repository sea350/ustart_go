package profile

import (
	"encoding/json"
	"html"
	"strings"

	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	uses "github.com/sea350/ustart_go/uses"
)

//TagStruct ... who knows at this point
type TagStruct struct {
	Tags []string
}

//AddTag ...
func AddTag(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	var ts TagStruct
	err := json.Unmarshal([]byte(r.FormValue("skillArray")), &ts.Tags)

	if err != nil {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Could not unmarshal")
		return
	}
	ID := session.Values["DocID"].(string)

	p := bluemonday.UGCPolicy()

	var validTags []string
	for t := range ts.Tags {
		ts.Tags[t] = p.Sanitize(ts.Tags[t])
		ts.Tags[t] = html.EscapeString(ts.Tags[t])
		isAllowed := uses.TagAllowed(client.Eclient, strings.ToLower(ts.Tags[t]))

		if isAllowed {
			validTags = append(validTags, ts.Tags[t])
		}
	}

	badgeTags, err := uses.BadgeSetupWID(client.Eclient, ID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	if len(badgeTags) > 0 {
		validTags = append(badgeTags, validTags...)
	}

	tagMap := make(map[string]int)
	for t := range validTags {
		tagMap[validTags[t]] = 0
	}
	validTags = nil
	for key := range tagMap {
		validTags = append(validTags, key)

	}

	err = post.UpdateUser(client.Eclient, ID, "Tags", validTags)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
