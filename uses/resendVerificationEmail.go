package uses

import (
	"log"
	"time"

	getUser "github.com/sea350/ustart_go/get/user"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//resendVerificationEmail ... resends User Verification Email
//Requires the user's email address
//Returns if the email failed to send
func ResendEmailVerification(eclient *elastic.Client, email string) {
	userID, err := getUser.UserIDByEmail(eclient, email)
	if err != nil {
		log.Println("Error uses/resendVerificationEmail line 15")
		log.Println(err)
		return
	}

	//Todo: Get user.type by email instead of by id
	user, err := getUser.UserByID(eclient, userID)
	if err != nil {
		log.Println("Error: uses/resendVerificationEmail line 24")
		log.Println(err)
		return
	}

	if !user.FirstLogin {
		token, err := GenerateRandomString(32)
		if err != nil {
			log.Println("Error: uses/resendVerificationEmail line 32")
			log.Println(err)
			return
		}

		err = postUser.UpdateUser(eclient, userID, "AuthenticationCode", token)
		if err != nil {
			log.Println("Error: uses/resendVerificationEmail line 39")
			log.Println(err)
			return
		}

		err = postUser.UpdateUser(eclient, userID, "AuthenticationCodeTime", time.Now())
		if err != nil {
			log.Println("Error: uses/resendVerificationEmail line 46")
			log.Println(err)
			return
		}

		subject := "Your verification link"
		link := "http://ustart.today:5002/ResetPassword/?email=" + email + "&verifCode=" + token
		r := NewRequest([]string{email}, subject)
		r.Send("/ustart/ustart_front/email_template.html", map[string]string{"username": user.Username, "link": link,
			"contentjuan":   "We received a request to reset your password for your Ustart Account. We would love to assist you!",
			"contentdos":    "Simply click the button below to verify your account",
			"contenttres":   "VERIFY ACCOUNT",
			"contentquatro": "a new account"})
	}
	return
}
