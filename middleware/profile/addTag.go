package profile

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//AddTag ...
func AddTag(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	ID := session.Values["DocID"].(string)

	tags := strings.Split(r.FormValue("skillArray"), `","`)
	if len(tags) > 0 {
		tags[0] = strings.Trim(tags[0], `["`)
		tags[len(tags)-1] = strings.Trim(tags[len(tags)-1], `"]`)
	}

	err = post.UpdateUser(client.Eclient, ID, "Tags", tags)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/addTag line 31")
	}
}
