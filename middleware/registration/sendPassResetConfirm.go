package registration

import (
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
)

//SendPassResetConfirm ... Renders a confirmation page
func SendPassResetConfirm(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	cs := client.ClientSide{}
	client.RenderSidebar(w, r, "templateNoUser2")
	client.RenderTemplate(w, r, "recoverpass-landing", cs)
}
