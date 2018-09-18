package project

import (
	"encoding/json"
	"html"
	"log"
	"net/http"
	"os"

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
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()

	ID := r.FormValue("skillArray")
	theID := r.FormValue("projectWidget")
	var ts []string
	err := json.Unmarshal([]byte(ID), &ts)

	if err != nil {
		log.Println("Could not unmarshal")
		return
	}

	for t := range ts {
		ts[t] = p.Sanitize(ts[t])
		ts[t] = html.EscapeString(ts[t])
	}

	err = post.UpdateProject(client.Eclient, theID, "Tags", ts)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
