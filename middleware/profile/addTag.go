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
		isAllowed := uses.TagAllowed(client.Eclient, ID, strings.ToLower(ts.Tags[t]))

		if isAllowed {
			validTags = append(validTags, ts.Tags[t])
		}
	}

	err = post.UpdateUser(client.Eclient, ID, "Tags", validTags)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
