package event

import (
	"net/http"

	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/event"
)

//UpdateEventTags ...
func UpdateEventTags(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := bluemonday.UGCPolicy()

	ID := r.FormValue("eventWidget")

	tags := strings.Split(p.Sanitize(r.FormValue("skillArray")), `","`)
	if len(tags) > 0 {
		tags[0] = strings.Trim(tags[0], `["`)
		tags[len(tags)-1] = strings.Trim(tags[len(tags)-1], `"]`)
	}

	err := post.UpdateEvent(client.Eclient, ID, "Tags", tags)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
