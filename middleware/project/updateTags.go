package project

import (
	"fmt"
	"net/http"
	"strings"

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
	}

	ID := r.FormValue("projectWidget")

	tags := strings.Split(r.FormValue("skillArray"), `","`)
	if len(tags) > 0 {
		tags[0] = strings.Trim(tags[0], `["`)
		tags[len(tags)-1] = strings.Trim(tags[len(tags)-1], `"]`)
	}

	err := post.UpdateProject(client.Eclient, ID, "Tags", tags)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error: middleware/project/updatetags line 31")
	}
}
