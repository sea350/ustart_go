package registration

import (
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/uses"
)

//ResendVerificationEmail ... Resends the email verification link to user
func ResendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	uses.SendVerificationEmail(client.Eclient, session.Values["Email"].(string))
	http.Redirect(w, r, "/registrationcomplete/", http.StatusFound)
}
