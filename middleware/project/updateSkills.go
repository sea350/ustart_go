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

	arrSkills := strings.Split(r.FormValue("skillArray"), `","`)
	if len(arrSkills) > 0 {
		arrSkills[0] = strings.Trim(arrSkills[0], `["`)
		arrSkills[len(arrSkills)-1] = strings.Trim(arrSkills[len(arrSkills)-1], `"]`)
	}

	p := bluemonday.UGCPolicy()
	var cleanSkills []string
	for idx := range arrSkills {
		cleanSkills = append(cleanSkills, p.Sanitize(arrSkills[idx]))
	}
	err := post.UpdateProject(client.Eclient, ID, "ListNeeded", cleanSkills)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
