package registration

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	getUser "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	bcrypt "golang.org/x/crypto/bcrypt"
)

//ResetPassword ... Reset's user's password
//Requires the user's email address
//Returns if the email failed to send
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		fmt.Println("Test1:", test1)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println("error:", err)
	}

	email := strings.ToLower(r.FormValue("email")) // we only client.Store lowercase emails in the db
	emailedToken := r.FormValue("verifCode")

	user, err := getUser.UserByEmail(client.Eclient, email)
	if err != nil {
		fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 34: User not found")
		fmt.Println(err)
		return
	}

	// If the time since authentication code was input is less than 1 hour
	if time.Since(user.AuthenticationCodeTime) < (time.Second * 3600) {
		if emailedToken == user.AuthenticationCode {
			newHashedPass, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 10)
			if err != nil {
				fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 40: Error generating password")
				fmt.Println(err)
				return
			}

			userID, err := getUser.UserIDByEmail(client.Eclient, email)
			if err != nil {
				fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 50: User not found")
				fmt.Println(err)
				return
			}

			err = post.UpdateUser(client.Eclient, userID, "Password", newHashedPass)
			if err != nil {
				fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 57: Error resetting password")
				fmt.Println(err)
				return
			}
			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", nil)
			if err != nil {
				fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 63: Unable to remove authentication code")
				fmt.Println(err)
			}

			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", nil)
			if err != nil {
				fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 69: Unable to remove authentication code time")
				fmt.Println(err)
			}

			http.Redirect(w, r, "/~", http.StatusFound)
		}

		fmt.Println("Success!")
	}
	cs := client.ClientSide{ErrorStatus: false}
	client.RenderSidebar(w, r, "templateNoUser2")
	client.RenderTemplate(w, r, "reset-new-pass", cs)
}
