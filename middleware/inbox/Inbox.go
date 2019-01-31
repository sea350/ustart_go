package inbox

import (
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Inbox ...
func Inbox(w http.ResponseWriter, r *http.Request) {
	// Session called session_please is retreived if it exists
	session, _ := client.Store.Get(r, "session_please")
	// check if docid exists within the session note: there is inconsistency with checking docid/username.
	test1, _ := session.Values["DocID"]
	// if no docid then redirect to main page
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	// if conditions of redirect are met, the next 4 lines do not execute
	// if docid exists then data is retreived from the uses.UserPage function and rendered accordingly
	userstruct, _, _, _, _ := uses.UserPage(client.Eclient, session.Values["Username"].(string), session.Values["DocID"].(string))
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string)}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderTemplate(w, r, "inbox-Nil", cs)
}
