package fail

import (
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
)

//Fail ... draws 404 page
func Fail(w http.ResponseWriter, r *http.Request) {
	cs := client.ClientSide{}
	session, _ := client.Store.Get(r, "session_please")
	// check DOCID instead
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		client.RenderSidebar(w, r, "template2-nil")
	} else {
		client.RenderSidebar(w, r, "templateNoUser2")
	}

	client.RenderTemplate(w, r, "404", cs)
}
