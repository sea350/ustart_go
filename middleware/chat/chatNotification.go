package chat

import (
	
	"net/http"

	"github.com/sea350/ustart_go/uses"

	"github.com/gorilla/websocket"
	"github.com/sea350/ustart_go/middleware/client"
)

var chatClients = make(map[string](map[*websocket.Conn]bool))
var chatBroadcast = make(chan chatNotif)

type chatNotif struct {
	UserID string `json:"UserID"`
	ChatID string `json:"ChatID"`
}

//HandleChatClients ...
func HandleChatClients(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	_, exists := chatClients[docID.(string)]
	if !exists {
		temp := make(map[*websocket.Conn]bool)
		temp[ws] = true
		chatClients[docID.(string)] = temp
	} else {
		temp := chatClients[docID.(string)]
		temp[ws] = true
		chatClients[docID.(string)] = temp
	}

	for {
		var notif chatNotif
		err := ws.ReadJSON(&notif)
		if err != nil {
			break
		}
	}

}

//HandleChatAlert ... deals with chat notifications, meant to be run on navbar
func HandleChatAlert() {
	for {
		// Grab the next message from the broadcast channel
		alert := <-chatBroadcast
		// Send it out to every client that is currently connected

		for clnt := range chatClients[alert.UserID] {
			head, err := uses.ConvertChatToFloatingHead(client.Eclient, alert.ChatID, alert.UserID)
			if err != nil {
				
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
				continue
			}

			err = clnt.WriteJSON(head)
			if err != nil {
				//log.Printf("error: %v", err)
				clnt.Close()
				delete(chatClients[alert.UserID], clnt)
			}
		}
	}
}
