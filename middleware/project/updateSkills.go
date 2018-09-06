package project

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"

	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
)

type SkillStruct struct {
	Skills []string
}

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

	var ss SkillStruct
	err := json.Unmarshal([]byte(skills), &ss.Skills)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		log.Println(err)
	}

	for s := range ss.Skills {
		ss.Skills[s] = p.Sanitize(ss.Skills[s])
	}
	err = post.UpdateProject(client.Eclient, ID, "ListNeeded", ss.Skills)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
