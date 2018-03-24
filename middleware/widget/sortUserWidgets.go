package widget

import (
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
)

//SortUserWidgets ... gets new array of widget ids from user page and updates user struct in ES
func SortUserWidgets(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	r.ParseForm()
	sortedWidgets := r.FormValue("sortedWidgets")

	fmt.Println(sortedWidgets)
}
