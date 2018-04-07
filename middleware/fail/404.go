package fail

import (
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
)

//Fail ... draws 404 page
func Fail(w http.ResponseWriter, r *http.Request) {
	cs := client.ClientSide{}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderTemplate(w, r, "404", cs)
}
