package registration

import (
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

func ResendEmailVerification(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	uses.SendVerificationEmail(client.Eclient, email)
}
