package project

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"

	"github.com/sea350/ustart_go/backend/middleware/client"
	post "github.com/sea350/ustart_go/backend/post/project"
)

//UpdateSkills ...
func UpdateSkills(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	ID := r.FormValue("projectWidget")
	skills := r.FormValue("skillArray")

	p := bluemonday.UGCPolicy()

	var ss []string
	err := json.Unmarshal([]byte(skills), &ss)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		log.Println(err)
	}

	for s := range ss {
		ss[s] = p.Sanitize(ss[s])
	}
	err = post.UpdateProject(client.Eclient, ID, "ListNeeded", ss)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
