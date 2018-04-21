package profile

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//AjaxChangeProjVisibility ... an ajax call that changes whether a project is visible on the user page
func AjaxChangeProjVisibility(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	ID, _ := session.Values["DocID"]
	if ID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	projID := r.FormValue("projectID")

	usr, err := get.UserByID(client.Eclient, ID.(string))
	if err != nil {
		fmt.Println(err)
		fmt.Println("err: middleware/profile/AjaxChangeProjVisibility line 27")
		return
	}

	for i := range usr.Projects {
		if usr.Projects[i].ProjectID == projID {
			if usr.Projects[i].Visible {
				usr.Projects[i].Visible = false
			} else {
				usr.Projects[i].Visible = true
			}
		}
	}

	err = post.UpdateUser(client.Eclient, ID.(string), "Projects", usr.Projects)
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an err: middleware/profile/addTag line 31")
	}
}