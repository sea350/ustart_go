package notification

import (
	
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/notification"
)

//MarkAsSeen ... wrestle with jimmy
//Designed for ajax
func MarkAsSeen(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	notifID := r.FormValue("notifID")
	if notifID == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"Blank notification ID")
		return
	}

	err := post.UpdateNotification(client.Eclient, notifID, "Seen", true)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
}
