package notification

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sea350/ustart_go/backend/middleware/client"
	properloading "github.com/sea350/ustart_go/backend/properloading"
	"github.com/sea350/ustart_go/backend/uses"
)

//AjaxScrollNotification ...
//Scrolls notifications
func AjaxScrollNotification(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	scrollID := r.FormValue("scrollID")

	sID, notifMap, _, err := properloading.ScrollNotifications(client.Eclient, docID.(string), scrollID)
	if err != nil && err != io.EOF {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)

	}
	var notifs []map[string]interface{}

	for notifID, notif := range notifMap {
		log.Println("NOTIF ID:", notifID)
		msg, url, err := uses.GenerateNotifMsgAndLink(client.Eclient, notif)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}

		notifAggregate := make(map[string]interface{})
		notifAggregate["ID"] = notifID
		notifAggregate["Data"] = notif
		notifAggregate["Message"] = msg
		notifAggregate["URL"] = url

		notifs = append(notifs, notifAggregate)
		// count++
		// if count == 5 {
		// 	break
		// }

	}

	sendData := make(map[string]interface{})
	sendData["notifications"] = notifs
	sendData["scrollID"] = sID

	// sendData["numUnread"] = proxy.NumUnread

	data, err := json.Marshal(sendData)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	log.Println(string(data))
	fmt.Fprintln(w, string(data))
}
