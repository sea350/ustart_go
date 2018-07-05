package project

import (
	"log"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//AjaxToggleFollow ... one click follow unfollow
func AjaxToggleFollow(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		return
	}

	projectID := r.FormValue("UNKOWN")

	err := uses.UserFollowProjectToggle(client.Eclient, test1.(string), projectID)
	if err != nil {
		log.Println("err: middleware/project/ajaxToggleFollow line 22")
		log.Println(err)
	}
}
