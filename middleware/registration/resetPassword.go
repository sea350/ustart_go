package registration

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	getUser "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
	bcrypt "golang.org/x/crypto/bcrypt"
	elastic "gopkg.in/olivere/elastic.v5"
)

var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

//ResetPassword ... Reset's user's password
//Requires the user's email address
//Returns if the email failed to send
func ResetPassword(eclient *elastic.Client, email string, w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	r.ParseForm()
	email := []byte(r.FormValue("email"))
	newPass := []byte(r.FormValue("newpass"))
	newHashedPass, err := bcrypt.GenerateFromPassword(newPass, 10)
	if err != nil {
		fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 28: Error generating password")
		fmt.Println(err)
		return err
	}

	userID, err := getUser.IDByUsername(eclient, email)
	if err != nil {
		fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 34: User not found")
		fmt.Println(err)
		return err
	}

	err = post.UpdateUser(eclient, userID, "Password", newHashedPass)
	if err != nil {
		fmt.Println("Error: /ustart_go/middleware/settings/resetPassword/ line 40: Error resetting password")
		fmt.Println(err)
		return err
	} else {
		fmt.Println("Success")
	}
	return err
}
