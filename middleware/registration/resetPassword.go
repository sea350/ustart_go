package registration

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	getUser "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	bcrypt "golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

//ResetPassword ... Reset's user's password
//Requires the user's email address
//Returns if the email failed to send
func ResetPassword(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	email := r.FormValue("email")
	email = strings.ToLower(email) // we only client.Store lowercase emails in the db
	newPass := []byte(r.FormValue("newpass"))
	newHashedPass, err := bcrypt.GenerateFromPassword(newPass, 10)
	if err != nil {
		fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 28: Error generating password")
		fmt.Println(err)
		return err
	}

	userID, err := getUser.IDByUsername(client.Eclient, email)
	if err != nil {
		fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 34: User not found")
		fmt.Println(err)
		return err
	}

	err = post.UpdateUser(client.Eclient, userID, "Password", newHashedPass)
	if err != nil {
		fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 40: Error resetting password")
		fmt.Println(err)
		return err
	} else {
		fmt.Println("Success")
	}
	return err
}
