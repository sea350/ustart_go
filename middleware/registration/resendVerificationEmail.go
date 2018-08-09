package registration

import (
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

func ResendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	fmt.Println(email)
	uses.SendVerificationEmail(client.Eclient, email)
}
