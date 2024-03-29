package registration

import (
	"errors"

	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"

	get "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
	uses "github.com/sea350/ustart_go/uses"
)

//SendPasswordResetEmail ... Sends password reset token link to user and saves token to their AuthenticationCode
//Requires a valid user email address
//Returns if there is an error
func SendPasswordResetEmail(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()

	email := p.Sanitize(r.FormValue("email"))

	var cs client.ClientSide

	//If the email isn't blank and it is in use...
	if email != "" {
		emailInUse, err := get.EmailInUse(client.Eclient, email)
		if err != nil {

			client.Logger.Println("Email: "+email+" | err: ", err)
		}

		if !emailInUse {
			cs = client.ClientSide{ErrorOutput: errors.New("No email found"), ErrorStatus: true}
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-forgot-pw", cs)
			return
		}
		token, err := uses.GenerateRandomString(32)
		if err != nil {

			client.Logger.Println("Email: "+email+" | err: ", err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-forgot-pw", cs)
			return
		}

		userID, err := get.UserIDByEmail(client.Eclient, email)
		if err != nil {

			client.Logger.Println("Email: "+email+" | err: ", err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-forgot-pw", cs)
			return
		}

		err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", time.Now())
		if err != nil {

			client.Logger.Println("DocID: "+userID+" | err: ", err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-forgot-pw", cs)
			return
		}

		err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", token)
		if err != nil {

			client.Logger.Println("DocID: "+userID+" | err: ", err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-forgot-pw", cs)
			return
		}

		user, err := get.UserByID(client.Eclient, userID)
		if err != nil {

			client.Logger.Println("DocID: "+userID+" | err: ", err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-forgot-pw", cs)
			return
		}

		subject := "Your password-reset link"
		link := globals.SiteURL + "/ResetPassword/?email=" + email + "&verifCode=" + token
		ree := uses.NewRequest([]string{email}, subject)
		ree.Send(
			globals.HTMLPATH+"email_template.html", map[string]string{
				"username":      user.Username,
				"link":          link,
				"contentjuan":   "We received a request to reset your password for your Ustart Account. We would love to assist you!",
				"contentdos":    "Simply click the button below to set a new password",
				"contenttres":   "CHANGE PASSWORD",
				"contentquatro": "a password reset"})

		http.Redirect(w, r, "/SendPassResetConfirm/", http.StatusFound)
		return
	}
	client.RenderSidebar(w, r, "templateNoUser2")
	client.RenderTemplate(w, r, "reset-forgot-pw", cs)
}
