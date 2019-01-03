package chat

import (
	"encoding/json"
	"fmt"
	
	"net/http"

	"github.com/sea350/ustart_go/uses"

	"github.com/sea350/ustart_go/middleware/client"
)

//AjaxNotificationLoad ... crawling in the 90s
//Designed for ajax
func AjaxNotificationLoad(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	heads, numUnread, err := uses.ChatAggregateNotifications(client.Eclient, docID.(string))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	sendData := make(map[string]interface{})
	sendData["numUnread"] = numUnread
	sendData["heads"] = heads

	data, err := json.Marshal(sendData)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	fmt.Fprintln(w, string(data))
}
