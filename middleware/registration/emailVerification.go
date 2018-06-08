package registration

import (
	"fmt"
	"net/http"
)

func EmailVerification(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	token := r.FormValue("verifCode")
	fmt.Println(email + "  " + token)
}
