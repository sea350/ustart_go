package registration

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

//ResetPassword ... Reset's user's password
//Requires the user's email address
//Returns if the email failed to send
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// session, _ := store.Get(r, "session_please")
	// test1, _ := session.Values["DocID"]
	// if test1 == nil {
	// 	fmt.Println(test1)
	// 	http.Redirect(w, r, "/~", http.StatusFound)
	// }

	// email := r.FormValue("email")
	// email = strings.ToLower(email) // we only client.Store lowercase emails in the db
	// newPass := []byte(r.FormValue("newpass"))
	// newHashedPass, err := bcrypt.GenerateFromPassword(newPass, 10)
	// if err != nil {
	// 	fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 28: Error generating password")
	// 	fmt.Println(err)
	// }

	// // If the time since authentication code was input is less than 1 hour
	// if time.Since(user.AuthenticationCodeTime) < (time.Second * 3600) {
	// 	if emailedToken == user.AuthenticationCode {
	// 		newHashedPass, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("newpass")), 10)
	// 		if err != nil {
	// 			fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 40: Error generating password")
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		userID, err := getUser.UserIDByEmail(client.Eclient, email)
	// 		if err != nil {
	// 			fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 50: User not found")
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 		err = post.UpdateUser(client.Eclient, userID, "Password", newHashedPass)
	// 		if err != nil {
	// 			fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 57: Error resetting password")
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", nil)
	// 		if err != nil {
	// 			fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 63: Unable to remove authentication code")
	// 			fmt.Println(err)
	// 		}

	// 		err = post.UpdateUser(client.Eclient, userID, "Password", newHashedPass)
	// 		if err != nil {
	// 			fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 40: Error resetting password")
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Println("Success")

	// 			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", nil)
	// 			if err != nil {
	// 				fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 69: Unable to remove authentication code time")
	// 				fmt.Println(err)
	// 			}
	// 		}

	// 	}
	// 	return
	// }
	return
}
