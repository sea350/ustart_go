package registration

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		log.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()

	email := p.Sanitize(r.FormValue("email"))

	var cs client.ClientSide

	//If the email isn't blank and it is in use...
	if email != "" {
		emailInUse, err := get.EmailInUse(client.Eclient, email)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}

		if !emailInUse {
			cs.ErrorOutput = errors.New("Invalid email")
			cs.ErrorStatus = true

			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reset-forgot-pw", cs)
			return
		} else {
			token, err := uses.GenerateRandomString(32)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				cs.ErrorStatus = true
				cs.ErrorOutput = err
				client.RenderSidebar(w, r, "templateNoUser2")
				client.RenderTemplate(w, r, "reset-forgot-pw", cs)
				return
			}

			userID, err := get.UserIDByEmail(client.Eclient, email)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				cs.ErrorStatus = true
				cs.ErrorOutput = err
				client.RenderSidebar(w, r, "templateNoUser2")
				client.RenderTemplate(w, r, "reset-forgot-pw", cs)
				return
			}

			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", time.Now())
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				cs.ErrorStatus = true
				cs.ErrorOutput = err
				client.RenderSidebar(w, r, "templateNoUser2")
				client.RenderTemplate(w, r, "reset-forgot-pw", cs)
				return
			}

			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", token)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				cs.ErrorStatus = true
				cs.ErrorOutput = err
				client.RenderSidebar(w, r, "templateNoUser2")
				client.RenderTemplate(w, r, "reset-forgot-pw", cs)
				return
			}

			user, err := get.UserByID(client.Eclient, userID)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				cs.ErrorStatus = true
				cs.ErrorOutput = err
				client.RenderSidebar(w, r, "templateNoUser2")
				client.RenderTemplate(w, r, "reset-forgot-pw", cs)
				return
			}

			subject := "Your password-reset link"
			link := "http://ustart.today:5002/ResetPassword/?email=" + email + "&verifCode=" + token
			r := uses.NewRequest([]string{email}, subject)
			r.Send(
				"/ustart/ustart_front/email_template.html", map[string]string{
					"username":      user.Username,
					"link":          link,
					"contentjuan":   "We received a request to reset your password for your Ustart Account. We would love to assist you!",
					"contentdos":    "Simply click the button below to set a new password",
					"contenttres":   "CHANGE PASSWORD",
					"contentquatro": "a password reset"})
			cs.Sent = "Email successfully sent!"
		}
	}
}
