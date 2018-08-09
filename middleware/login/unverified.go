package login

import (
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
)

//Unverified ... there's no place like it
func Unverified(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	}
	session.Save(r, w)
	cs := client.ClientSide{}
	client.RenderSidebar(w, r, "templateNoUser2")
	client.RenderTemplate(w, r, "resend-email", cs)
}
