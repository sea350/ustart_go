package project

import (
	"log"
	"net/http"
	"os"
	"strings"

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

	ID := r.FormValue("projectWidget")

	tags := strings.Split(ID, `","`)

	if len(tags) > 0 {
		tags[0] = strings.Trim(tags[0], `["`)
		tags[len(tags)-1] = strings.Trim(tags[len(tags)-1], `"]`)
	}

	var cleanTags []string
	for idx := range tags {
		cleanTags = append(cleanTags, p.Sanitize(tags[idx]))
	}

	err := post.UpdateProject(client.Eclient, ID, "Tags", cleanTags)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
