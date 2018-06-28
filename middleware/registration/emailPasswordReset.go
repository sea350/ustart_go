package registration

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/sea350/ustart_go/middleware/client"

	get "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
	uses "github.com/sea350/ustart_go/uses"
)

//SendPasswordResetEmail ... Sends password reset token link to user and saves token to their AuthenticationCode
//Requires a valid user email address
//Returns if there is an error
func SendPasswordResetEmail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	var cs client.ClientSide

	defer client.RenderSidebar(w, r, "templateNoUser2")
	defer client.RenderTemplate(w, r, "reset-forgot-pw", cs)

	//If the email isn't blank and it is in use...
	if email != "" {
		emailInUse, err := get.EmailInUse(client.Eclient, email)
		if err != nil {
			fmt.Println("Error: ustart_go/middleware/registration/emailPasswordReset Line 21")
			fmt.Println(err)
		}

		if emailInUse {
			token, err := uses.GenerateRandomString(32)
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 16: Error generating token")
				fmt.Println(err)
				return
			}

			userID, err := get.UserIDByEmail(client.Eclient, email)
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 16: Unable to retreive userID by email")
				fmt.Println(err)
				return
			}

			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", time.Now())
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 16: Error posting user")
				fmt.Println(err)
				return
			}

			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", token)
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 16: Error posting user")
				fmt.Println(err)
				return
			}

			from := "ustarttestemail@gmail.com"
			pass := "Ust@rt20!8~~"
			body := "http://ustart.today:5002/ResetPassword/?email=" + email + "&verifCode=" + token
			msg := "From: " + from + "\n" + "To: " + email + "\n" + "Subject: UStart Password Reset\n\n" + body

			err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{email}, []byte(msg))
			if err != nil {
				log.Printf("smtp error: %s", err)
				return
			}

		}
	}
}
