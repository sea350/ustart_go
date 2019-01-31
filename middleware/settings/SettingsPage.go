package settings

import (
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Settings ...
func Settings(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	userstruct, _, _, _, _ := uses.UserPage(client.Eclient, session.Values["Username"].(string), session.Values["DocID"].(string))
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string)}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderTemplate(w, r, "settings-Nil", cs)

}
