package chat

import (
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
)

//Page ... draws chat page
func Page(w http.ResponseWriter, r *http.Request) {
	cs := client.ClientSide{}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderTemplate(w, r, "cuzsteventoldmeto", cs)
	//chatID := r.URL.Path[4:]

	// var h = hubAlt{
	// 	broadcast:  make(chan messageAlt),
	// 	register:   make(chan subscription),
	// 	unregister: make(chan subscription),
	// 	rooms:      make(map[string]map[*connection]bool),
	// }
	go handleMessages()
	//go ServeWs(w, r)
	// go h.Run()
}
