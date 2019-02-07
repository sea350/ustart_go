package project

import (
	"encoding/json"
	"html"

	"github.com/sea350/ustart_go/uses"

	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
)

//UpdateTags ...
func UpdateTags(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()

	ID := r.FormValue("skillArray")
	theID := r.FormValue("projectWidget")
	var ts []string
	err := json.Unmarshal([]byte(ID), &ts)

	if err != nil {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Could not unmarshal")
		return
	}

	var updatedTags []string
	for t := range ts {
		ts[t] = p.Sanitize(ts[t])
		ts[t] = html.EscapeString(ts[t])
		allowed := uses.TagAllowed(client.Eclient, ts[t])
		if allowed {
			updatedTags = append(updatedTags, ts[t])
		}

	}

	err = post.UpdateProject(client.Eclient, theID, "Tags", updatedTags)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
