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
	go handleMessages()
}
