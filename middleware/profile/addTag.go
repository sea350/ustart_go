package profile

import (
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
)

//AddTag ...
func AddTag(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	//ID := session.Values["DocID"].(string)

	/*
		usr, err := get.UserByID(client.Eclient, ID)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an err: middleware/profile/addTag line 25")
		}
	*/
	fmt.Println(w, r.FormValue("skillArray"))
	/*
		usr.Tags = append(usr.Tags, r.FormValue("skillArray"))

		err = post.UpdateUser(client.Eclient, ID, "Tags", usr.Tags)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an err: middleware/profile/addTag line 31")
		}
	*/
	fmt.Fprintln(w, "minhazaur")
}
