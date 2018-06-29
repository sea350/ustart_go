package chat

import (
	"fmt"
	"sync"
	"time"
)

//Hub ... what's all the hubbub bub?
type Hub struct {
	// the mutex to protect connections
	connectionsMx sync.RWMutex

	// Registered connections.
	connections map[*Connection]struct{}

	// Inbound messages from the connections.
	broadcast chan []byte

	logMx sync.RWMutex
	log   [][]byte
}

//NewHub ... makes new hub
func NewHub() *Hub {
	h := &Hub{
		connectionsMx: sync.RWMutex{},
		broadcast:     make(chan []byte),
		connections:   make(map[*Connection]struct{}),
	}

	go func() {
		for {
			msg := <-h.broadcast
			h.connectionsMx.RLock()
			for c := range h.connections {
				select {
				case c.send <- msg:
				// stop trying to send to this connection after trying for 1 second.
				// if we have to stop, it means that a reader died so remove the connection also.
				case <-time.After(1 * time.Second):
					//fmt.Printf("shutting down connection %s \n", c)
					fmt.Printf("kill me")
					h.removeConnection(c)
				}
			}
			h.connectionsMx.RUnlock()
		}
	}()
	return h
}

func (h *Hub) addConnection(conn *Connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	h.connections[conn] = struct{}{}
}

func (h *Hub) removeConnection(conn *Connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.send)
	}
}
