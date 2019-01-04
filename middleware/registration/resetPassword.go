package registration

import (
	"errors"

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
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()

	cs := client.ClientSide{}

	email := strings.ToLower(r.FormValue("email")) // we only client.Store lowercase emails in the db
	emailedToken := r.FormValue("verifCode")

	user, err := getUser.UserByEmail(client.Eclient, email)
	if err != nil {

		client.Logger.Println("Email: "+email+" | err: %s", err)
		cs.ErrorOutput = errors.New("A problem has occurred")
		cs.ErrorStatus = true
		client.RenderSidebar(w, r, "templateNoUser2")
		client.RenderTemplate(w, r, "reset-new-pass", cs)
		return
	}

	// If the time since authentication code was input is less than 1 hour
	if time.Since(user.AuthenticationCodeTime) < (time.Second*3600) && emailedToken == user.AuthenticationCode && r.FormValue("password") != `` {
		newHashedPass, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 10)
		if err != nil {

			client.Logger.Println("Email: "+email+" | err: %s", err)
			cs.ErrorOutput = errors.New("A problem has occurred")
			cs.ErrorStatus = true
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-new-pass", cs)
			return
		}

		userID, err := getUser.UserIDByEmail(client.Eclient, email)
		if err != nil {

			client.Logger.Println("Email: "+email+" | err: %s", err)
			cs.ErrorOutput = errors.New("A problem has occurred")
			cs.ErrorStatus = true
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-new-pass", cs)
			return
		}

		err = post.UpdateUser(client.Eclient, userID, "Password", newHashedPass)
		if err != nil {

			client.Logger.Println("DocID: "+userID+" | err: %s", err)
			cs.ErrorOutput = errors.New("A problem has occurred")
			cs.ErrorStatus = true
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-new-pass", cs)
			return
		}

		err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", nil)
		if err != nil {

			client.Logger.Println("DocID: "+userID+" | err: %s", err)
		}

		err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", nil)
		if err != nil {

			client.Logger.Println("DocID: "+userID+" | err: %s", err)
		}

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		cs.ErrorOutput = errors.New("Authentication token invalid")
		cs.ErrorStatus = true
	}

	client.RenderSidebar(w, r, "templateNoUser2")
	client.RenderTemplate(w, r, "reset-new-pass", cs)
}
