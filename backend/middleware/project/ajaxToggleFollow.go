package project

import (
	"log"
	"net/http"
	"os"

	"github.com/sea350/ustart_go/backend/uses"

	client "github.com/sea350/ustart_go/backend/middleware/client"
)

//AjaxToggleFollow ... one click follow unfollow
func AjaxToggleFollow(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		return
	}

	projectID := r.FormValue("projectID")

	err := uses.UserFollowProjectToggle(client.Eclient, test1.(string), projectID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
