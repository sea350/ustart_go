package settings

import (
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//AjaxToggleUserActive ... toggle's the user's "Status" field
func AjaxToggleUserActive(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	status := r.FormValue("work_availability")
	if status == `` {
		return
	}

	var statusBool bool
	if status == `on` {
		statusBool = true
	} else if status == `off` {
		statusBool = false
	} else {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | err: Improper argument passed in")

	}

	err := post.UpdateUser(client.Eclient, test1.(string), "Status", statusBool)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

}
