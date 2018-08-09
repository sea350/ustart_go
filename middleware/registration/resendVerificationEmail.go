package registration

import (
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
)

//ResendVerificationEmail ... Resends the email verification link to user
func ResendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	fmt.Println(session.Values["Email"])
	//uses.SendVerificationEmail(client.Eclient, email)
}
