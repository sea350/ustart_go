package fail

import (
	"net/http"

	"src/github.com/sea350/ustart_go/middleware/client"
)

func fail(w http.ResponseWriter, r *http.Request) {
	cs := client.ClientSide{}
	client.RenderTemplate(w, "template2-nil", cs)
	client.RenderTemplate(w, "404", cs)
}
