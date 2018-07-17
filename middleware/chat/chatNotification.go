package chat

import (
	"github.com/gorilla/websocket"
)

var chatNotification = make(map[string](*websocket.Conn))
