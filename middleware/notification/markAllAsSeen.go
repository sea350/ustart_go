package notification

import (
	"net/http"

	get "github.com/sea350/ustart_go/get/notification"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/notification"
)

//MarkAllAsSeen ... wrestle with jimmy
//Designed for ajax
func MarkAllAsSeen(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	ids, err := get.UnseenNotificationIDsByUserID(client.Eclient, docID.(string))
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	for _, id := range ids {
		err := post.UpdateNotification(client.Eclient, id, "Seen", true)
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
	}
}
