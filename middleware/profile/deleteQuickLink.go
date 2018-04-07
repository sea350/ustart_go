package profile

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
)

//DeleteQuickLink ...
func DeleteQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	username := test1.(string)
	ID := session.Values["DocID"].(string)

	usr, err := get.UserByID(client.Eclient, ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/deleteQuickLink line 25")
	}

	deleteTitle := r.FormValue("userLinkDesc")
	deleteURL := r.FormValue("userLink")

	var newArr []types.Link

	if len(usr.QuickLinks) <= 1 {
		err := post.UpdateUser(client.Eclient, ID, "QuickLinks", newArr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an err: middleware/profile/deleteQuickLink line 39")
		}
		http.Redirect(w, r, "/profile/"+username, http.StatusFound)
		return
	}

	target := -1
	for index, link := range usr.QuickLinks {

		if link.Name == deleteTitle && link.URL == deleteURL {
			target = index
			fmt.Println(target)
			break
		}
	}

	if target == -1 {
		fmt.Println("deleted object not found")
		fmt.Println("this is an err, middleware/profile/deleteQuickLink line 57")
		newArr = usr.QuickLinks
	} else if (target + 1) < len(usr.QuickLinks) {
		newArr = append(usr.QuickLinks[:target], usr.QuickLinks[(target+1):]...)
	} else {
		newArr = usr.QuickLinks[:target]
	}

	err = post.UpdateUser(client.Eclient, ID, "QuickLinks", newArr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/deleteQuickLink line 68")
	}

	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}
