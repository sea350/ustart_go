package chat

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sea350/ustart_go/middleware/client"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var chatroom = make(map[string](map[*websocket.Conn]bool))
var broadcast = make(chan carrierMessage) // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//DO NOT EXPORT
type carrierMessage struct {
	Username  string    `json:"Username"`
	DocID     string    `json:"DocID"`
	Message   string    `json:"Message"`
	ChatID    string    `json:"ChatID"`
	TimeStamp time.Time `json:"TimeStamp"`
}

//HandleConnections ...
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	chatID := r.URL.Path[4:]
	// Register our new client
	_, exists := chatroom[chatID]
	if !exists {
		temp := make(map[*websocket.Conn]bool)
		temp[ws] = true
		chatroom[chatID] = temp
	} else {
		temp := chatroom[chatID]
		temp[ws] = true
		chatroom[chatID] = temp
	}

	for {
		var msg carrierMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		msg.ChatID = chatID
		msg.DocID = docID.(string)
		msg.TimeStamp = time.Now()
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		/*
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, "debug text")
			log.Println("channel #" + msg.ChatID)
			log.Printf("message: %v \n", msg)
			log.Println(chatroom[msg.ChatID])
		*/
		for client := range chatroom[msg.ChatID] {
			err := client.WriteJSON(msg)
			if err != nil {
				//log.Printf("error: %v", err)
				client.Close()
				delete(chatroom[msg.ChatID], client)
			}
		}
	}
}
