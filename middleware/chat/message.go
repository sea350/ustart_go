package chat

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
var channels = make(map[string](chan Message))

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
	clients[ws] = true

	fmt.Println("debug text: middleware/chat/message line 41")
	fmt.Println(chatID)

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
		fmt.Println("debug text: middleware/chat/message line 55")
		fmt.Println(msg)
		channels[chatID] <- msg
		fmt.Println("message sent")
	}
}

func handleMessages(chatID string) {
	for {
		// Grab the next message from the broadcast channel
		msg := <-channels[chatID]
		// Send it out to every client that is currently connected
		fmt.Println("debug text: middleware/chat/message line 67")
		fmt.Println(msg)
		fmt.Println("message received")
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
