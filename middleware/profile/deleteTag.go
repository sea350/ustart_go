package profile

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//DeleteTag ...
func DeleteTag(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	ID := session.Values["DocID"].(string)

	usr, err := get.UserByID(client.Eclient, ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/deleteTag line 25")
	}

	deleteTag := r.FormValue("UNKNOWN")

	var newArr []string

	if len(usr.Tags) == 1 {
		err := post.UpdateUser(client.Eclient, ID, "Tags", newArr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an err: middleware/profile/deleteTag line 35")
		}
		return
	}

	target := -1
	for index, tag := range usr.Tags {

		if tag == deleteTag {
			target = index
			fmt.Println(target)
			break
		}
	}

	if target == -1 {
		fmt.Println("deleted object not found")
		fmt.Println("this is an err, middleware/profile/deleteTag line 54")
		newArr = usr.Tags
	} else if (target + 1) < len(usr.Tags) {
		newArr = append(usr.Tags[:target], usr.Tags[(target+1):]...)
	} else {
		newArr = usr.Tags[:target]
	}

	err = post.UpdateUser(client.Eclient, ID, "Tags", newArr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/deleteTag line 31")
	}
}
