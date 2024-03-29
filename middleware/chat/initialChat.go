package chat

import (
	"encoding/json"
	"fmt"
	
	"net/http"

	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"

	getChat "github.com/sea350/ustart_go/get/chat"
	"github.com/sea350/ustart_go/middleware/client"
)

//InitialChat ... crawling in the 90s
//Designed for ajax
func InitialChat(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	chatURL := r.FormValue("chatUrl")

	agg := make(map[string]interface{})
	//agg["eavesdroppers"] = make(map[string]types.FloatingHead)

	valid, actualChatID, otherUsr, err := uses.ChatVerifyURL(client.Eclient, chatURL, docID.(string))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	if !valid {
		return
	}

	if actualChatID != `` {
		idx, msgs, err := uses.ChatLoad(client.Eclient, actualChatID, -1, 50)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			return
		}
		agg["Messages"] = msgs
		agg["Index"] = idx

		chat, err := getChat.ConvoByID(client.Eclient, actualChatID)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			return
		}

		temp := make(map[string]types.FloatingHead)

		for idx := range chat.Eavesdroppers {
			head, err := uses.ConvertUserToFloatingHead(client.Eclient, chat.Eavesdroppers[idx].DocID)
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}

			temp[chat.Eavesdroppers[idx].DocID] = head
			agg["Eavesdroppers"] = temp
		}

	} else {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, otherUsr)
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		temp := make(map[string]types.FloatingHead)
		temp[otherUsr] = head

		head, err = uses.ConvertUserToFloatingHead(client.Eclient, docID.(string))
		if err != nil {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
		temp[docID.(string)] = head
		agg["Eavesdroppers"] = temp
	}

	data, err := json.Marshal(agg)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	fmt.Fprintln(w, string(data))
}
