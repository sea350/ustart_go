package chat

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var chatroom = make(map[string](map[*websocket.Conn]bool))
var broadcast = make(chan Message) // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//Message ... Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
	ChatID   string `json:"chatID"`
}

//HandleConnections ...
func HandleConnections(w http.ResponseWriter, r *http.Request) {
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
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		fmt.Println("debug text: middleware/chat/message line 67")
		fmt.Println("channel #" + msg.ChatID)
		fmt.Printf("message: %v \n", msg)
		fmt.Println(chatroom[msg.ChatID])

		for client := range chatroom[msg.ChatID] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(chatroom[msg.ChatID], client)
			}
		}
	}
}
