package registration

import (
	"fmt"
	"net/http"
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
			fmt.Println("Error: ustart_go/middleware/registration/emailPasswordReset Line 30: Unable to retrieve email")
			fmt.Println(err)
		}

		if emailInUse {
			token, err := uses.GenerateRandomString(32)
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 37: Error generating token")
				fmt.Println(err)
				return
			}

			userID, err := get.UserIDByEmail(client.Eclient, email)
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 44: Unable to retreive userID by email")
				fmt.Println(err)
				return
			}

			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", time.Now())
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 51: Error posting user")
				fmt.Println(err)
				return
			}

			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", token)
			if err != nil {
				fmt.Println("Error ustart_go/middleware/registration/emailPasswordReset line 58: Error posting user")
				fmt.Println(err)
				return
			}

			subject := "Your verification link"
			link := "http://ustart.today:5002/ResetPassword/?email=" + email + "&verifCode=" + token
			r := uses.NewRequest([]string{email}, subject)
			r.Send("/ustart/ustart_front/email_template.html", map[string]string{"username": email, "link": link})
		}
	}
}
