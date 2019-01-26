package notification

import (
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/notification"
)

//RemoveNotification ... wrestle with jimmy
//Designed for ajax
func RemoveNotification(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	notifID := r.FormValue("notifID")
	if notifID == `` {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Blank notification ID")
		return
	}

	err := post.UpdateNotification(client.Eclient, notifID, "Invisible", true)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
