package chat

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var chatroom = make(map[string](map[*websocket.Conn]string))
var broadcast = make(chan types.Message) // broadcast channel
var convoLock sync.Mutex

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	if !valid {
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
	} else if actualChatID != `` {
		_, exists := chatroom[actualChatID]
		if !exists {
			temp := make(map[*websocket.Conn]string)
			temp[ws] = docID.(string)
			chatroom[actualChatID] = temp
		} else {
			temp := chatroom[actualChatID]
			temp[ws] = docID.(string)
			chatroom[actualChatID] = temp
		}
		err = postChat.MarkAsRead(client.Eclient, docID.(string), actualChatID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}

	for {
		var msg types.Message
		var notif chatNotif
		notifyThese := []string{}
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(chatroom[actualChatID], ws)
			break
		}
		if len(msg.Content) > 500 {
			continue
		}

		msg = types.Message{SenderID: docID.(string), TimeStamp: time.Now(), Content: msg.Content, ConversationID: actualChatID}
		if actualChatID == `` && chatURL != `` {
			newConvoID, err := uses.ChatFirst(client.Eclient, msg, docID.(string), dmTargetUserID)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
			notifyThese = append(notifyThese, dmTargetUserID)
			notifyThese = append(notifyThese, docID.(string))
			actualChatID = newConvoID
			temp := make(map[*websocket.Conn]string)
			temp[ws] = docID.(string)
			chatroom[actualChatID] = temp
			msg.ConversationID = actualChatID

		} else if actualChatID != `` && chatURL != `` {
			notifyThese, err = uses.ChatSend(client.Eclient, msg)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
		}

		// Send the newly received message to the broadcast channel
		broadcast <- msg

		//send notification here
		notif.ChatID = actualChatID
		for _, id := range notifyThese {
			notif.UserID = id
			chatBroadcast <- notif
		}
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		for clnt, docID := range chatroom[msg.ConversationID] {
			err := clnt.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				clnt.Close()
				delete(chatroom[msg.ConversationID], clnt)
				return
			}
			err = postChat.MarkAsRead(client.Eclient, docID, msg.ConversationID)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}

		}
	}
}
