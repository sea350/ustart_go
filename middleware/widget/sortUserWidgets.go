package widget

import (
	"log"
	"net/http"
	"strings"

	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//SortUserWidgets ... gets new array of widget ids from user page and updates user struct in ES
func SortUserWidgets(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	sortedWidgets := r.FormValue("sortedWidgets")
	if r.FormValue("pageID") != session.Values["DocID"].(string) {
		return
	}

	ids := strings.Split(sortedWidgets, `","`)
	if len(ids) > 0 {
		ids[0] = strings.Trim(ids[0], `["`)
		ids[len(ids)-1] = strings.Trim(ids[len(ids)-1], `"]`)
	}

	err := post.UpdateUser(client.Eclient, session.Values["DocID"].(string), "UserWidgets", ids)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
}
