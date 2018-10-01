package login

import (
	"net/http"

	client "github.com/sea350/ustart_go/backend/middleware/client"
)

//Logout ... lougout
func Logout(w http.ResponseWriter, r *http.Request) {
	// Session called session_please is retreived if it exists
	session, _ := client.Store.Get(r, "session_please")
	// check if docid exists within the session note: there is inconsistency with checking docid/username.
	test1, _ := session.Values["DocID"]
	if test1 != nil { // if logged in
		session.Options.MaxAge = -1 // kills session
		session.Save(r, w)          // save changes to session
		http.Redirect(w, r, "/~", http.StatusFound)
		return // bring back to homepage
	}
}
