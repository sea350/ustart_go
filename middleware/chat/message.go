package chat

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

type room struct {
	sockets map[*websocket.Conn]string
	lock    sync.Mutex
}

//var clients = make(map[*websocket.Conn]bool) // connected clients
var chatroom = make(map[string]*room)
var broadcast = make(chan types.Message) // broadcast channel
var startChatLocks = make(map[string]*sync.Mutex)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func sendAndNotify(msg types.Message, notif chatNotif, notify []string) {
	chatroom[msg.ConversationID].lock.Lock()
	defer chatroom[msg.ConversationID].lock.Unlock()
	// Send the newly received message to the broadcast channel
	broadcast <- msg

	//send notification here

	for _, id := range notify {
		notif.UserID = id
		chatBroadcast <- notif
	}
}

//HandleConnections ...
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	chatURL := r.URL.Path[4:]

	//security checks before socket is opened
	valid, actualChatID, dmTargetUserID, err := uses.ChatVerifyURL(client.Eclient, chatURL, docID.(string))
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	if !valid {
		return
	}

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		client.Logger.Panic(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	_, registered := startChatLocks[docID.(string)]
	if !registered {
		startChatLocks[docID.(string)] = &sync.Mutex{}
	}

	/*IF YOU WAND GLOBAL CHAT ENABLED, DO THIS
		if chatURL == `` {
		_, exists := chatroom[``]
		if !exists {
			temp := make(map[*websocket.Conn]string)
			temp[ws] = docID.(string)
			chatroom[``] = temp
		} else {
			temp := chatroom[``]
			temp[ws] = docID.(string)
			chatroom[``] = temp
		}
	} else
	*/

	// Register our new client
	if actualChatID != `` {
		_, exists := chatroom[actualChatID]
		if !exists {
			temp := make(map[*websocket.Conn]string)
			temp[ws] = docID.(string)

			chatroom[actualChatID] = &room{sockets: temp}
		} else {
			temp := chatroom[actualChatID].sockets
			temp[ws] = docID.(string)
			chatroom[actualChatID] = &room{sockets: temp}
		}
		err = postChat.MarkAsRead(client.Eclient, docID.(string), actualChatID)
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
	}

	for {
		var msg types.Message
		var notif chatNotif
		notifyThese := []string{}
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			//client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			//DONT REPORT THIS ERROR
			_, exists1 := chatroom[actualChatID]
			if exists1 {
				_, exists2 := chatroom[actualChatID].sockets[ws]
				if exists2 {
					delete(chatroom[actualChatID].sockets, ws)
				}

			}
			_, exists3 := startChatLocks[docID.(string)]
			if exists3 {
				delete(startChatLocks, docID.(string))
			}
			break
		}

		if actualChatID == `` && chatURL != `` { // this needs to be done outside because the variables are liable to change
			startChatLocks[docID.(string)].Lock()
		}
		if len(msg.Content) > 500 {
			continue
		}

		msg = types.Message{SenderID: docID.(string), TimeStamp: time.Now(), Content: msg.Content, ConversationID: actualChatID}
		if actualChatID == `` && chatURL != `` {
			client.Logger.Println("Debug text: chat ID = " + actualChatID + " | chatUrl = " + chatURL)
			newConvoID, err := uses.ChatFirst(client.Eclient, msg, docID.(string), dmTargetUserID)
			if err != nil {
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
				startChatLocks[docID.(string)].Unlock()
				continue
			}
			notifyThese = append(notifyThese, dmTargetUserID)
			notifyThese = append(notifyThese, docID.(string))
			actualChatID = newConvoID
			temp := make(map[*websocket.Conn]string)
			temp[ws] = docID.(string)
			chatroom[actualChatID] = &room{sockets: temp}
			msg.ConversationID = actualChatID
			startChatLocks[docID.(string)].Unlock()

		} else if actualChatID != `` && chatURL != `` {
			notifyThese, err = uses.ChatSend(client.Eclient, msg)
			if err != nil {
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
				continue
			}
		} else {
			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | err: Unexpected condition met")
			continue
		}
		notif.ChatID = actualChatID
		sendAndNotify(msg, notif, notifyThese)
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		for clnt, docID := range chatroom[msg.ConversationID].sockets {

			err := clnt.WriteJSON(msg)
			if err != nil {

				client.Logger.Printf("error: %v", err)
				clnt.Close()
				delete(chatroom[msg.ConversationID].sockets, clnt)
				continue
			}
			err = postChat.MarkAsRead(client.Eclient, docID, msg.ConversationID)
			if err != nil {

				client.Logger.Println("DocID: "+docID+" | err: ", err)
			}

		}

	}
}
