package profile

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
)

//AddQuickLink ...
func AddQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	ID := session.Values["DocID"].(string)

	usr, err := get.UserByID(client.Eclient, ID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/addQuickLink line 25")
	}

	usr.QuickLinks = append(usr.QuickLinks, types.Link{Name: r.FormValue("userLinkDesc"), URL: r.FormValue("userLink")})

	err = post.UpdateUser(client.Eclient, ID, "QuickLinks", usr.QuickLinks)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/addQuickLink line 31")
	}

	fmt.Fprint(w, r.FormValue("userLinkDesc"))
	fmt.Fprint(w, r.FormValue("userLink"))

}
