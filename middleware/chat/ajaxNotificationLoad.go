package chat

import (
	"encoding/json"
	"fmt"
	"strconv"

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

	var i int
	index := r.FormValue("index")
	if index == `` {
		i = 0
	} else {
		integer, err := strconv.Atoi(index)
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			return
		}
		i = integer
	}

	heads, numUnread, err := uses.ChatAggregateNotifications(client.Eclient, docID.(string), i)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	sendData := make(map[string]interface{})
	sendData["numUnread"] = numUnread
	sendData["heads"] = heads

	data, err := json.Marshal(sendData)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	fmt.Fprintln(w, string(data))
}
