package login

import (
	"net/http"

	client "github.com/sea350/ustart_go/backend/middleware/client"
)

/* The following 2 lines do not need to be repeated in each subfolder. They can be separated just like everything else.
The first represents the connection to the eS cluster. The second corresponds to a session store. This is not a proper
way to handle the session store. */

//Home ... there's no place like it
func Home(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	}
	session.Save(r, w)
	cs := client.ClientSide{}
	client.RenderTemplate(w, r, "templateNoUser2", cs)
	client.RenderTemplate(w, r, "nil-index2", cs)
	client.RenderTemplate(w, r, "template-footer-nil", cs)
}
