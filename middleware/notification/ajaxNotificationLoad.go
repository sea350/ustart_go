package notification

import (
	"encoding/json"
	"fmt"
	
	"net/http"

	get "github.com/sea350/ustart_go/get/notification"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/notification"
	"github.com/sea350/ustart_go/uses"
)

//AjaxNotificationLoad ... crawling in the 90s
//Designed for ajax
func AjaxNotificationLoad(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	var notifs []map[string]interface{}

	id, proxy, err := get.ProxyNotificationByUserID(client.Eclient, docID.(string))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	err = post.ResetUnseen(client.Eclient, id)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	count := 0
	for _, id := range proxy.NotificationCache {
		notif, err := get.NotificationByID(client.Eclient, id)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			continue
		}
		if notif.Invisible {
			continue
		}

		msg, url, err := uses.GenerateNotifMsgAndLink(client.Eclient, notif)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			continue
		}

		notifAggregate := make(map[string]interface{})
		notifAggregate["ID"] = id
		notifAggregate["Data"] = notif
		notifAggregate["Message"] = msg
		notifAggregate["URL"] = url
		notifs = append(notifs, notifAggregate)
		count++
		if count == 5 {
			break
		}

	}

	sendData := make(map[string]interface{})
	sendData["notifications"] = notifs
	sendData["numUnread"] = proxy.NumUnseen

	data, err := json.Marshal(sendData)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	fmt.Fprintln(w, string(data))
}
