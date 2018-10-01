package notification

import (
	"log"
	"net/http"

	"github.com/sea350/ustart_go/backend/middleware/client"
	post "github.com/sea350/ustart_go/backend/post/notification"
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Blank notification ID")
		return
	}

	err := post.UpdateNotification(client.Eclient, notifID, "Invisible", true)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
}
