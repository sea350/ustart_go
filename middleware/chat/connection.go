package chat

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

//Connection ... it feels so strong between us
type Connection struct {
	// Buffered channel of outbound messages.
	send chan []byte

	// The hub.
	h *Hub
}

func (c *Connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			break
		}
		c.h.broadcast <- message
	}
}

func (c *Connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for message := range c.send {
		err := wsConn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

//WsHandler ... such strong hands
type WsHandler struct {
	H *Hub
}

func (wsh WsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("error upgrading %s", err)
		return
	}
	c := &Connection{send: make(chan []byte, 256), h: wsh.H}
	c.h.addConnection(c)
	defer c.h.removeConnection(c)
	var wg sync.WaitGroup
	wg.Add(2)
	go c.writer(&wg, wsConn)
	go c.reader(&wg, wsConn)
	wg.Wait()
	wsConn.Close()
}
