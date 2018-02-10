package profile

import (
	"fmt"
	"net/http"
	"strings"

	get "github.com/sea350/ustart_go/get/user"
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
	}

	ID := session.Values["DocID"].(string)
	tags := strings.Split(r.FormValue("skillArray"), `","`)
	tags[0] = strings.Trim(tags[1], `["`)
	tags[len(tags)-1] = strings.Trim(tags[1], `"]`)
	usr, err := get.UserByID(client.Eclient, ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/addTag line 25")
	}

	usr.Tags = append(usr.Tags)

	err = post.UpdateUser(client.Eclient, ID, "Tags", usr.Tags)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/addTag line 31")
	}

	fmt.Fprintln(w, "minhazaur")
}
