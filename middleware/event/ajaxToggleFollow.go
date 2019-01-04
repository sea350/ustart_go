package event

import (
	
	"net/http"
	

	"github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//AjaxEventToggleFollow ... one click follow unfollow
func AjaxEventToggleFollow(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		return
	}

	eventID := r.FormValue("eventID") // eventID

	err := uses.UserFollowEventToggle(client.Eclient, test1.(string), eventID)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
