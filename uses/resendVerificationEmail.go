package uses

import (
	"fmt"
	"log"
	"net/smtp"

	getUser "github.com/sea350/ustart_go/get/user"
	postUser "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//resendVerificationEmail ... resends User Verification Email
//Requires the user's email address
//Returns if the email failed to send
func ResendEmailVerification(eclient *elastic.Client, email string) {
	token, err := GenerateRandomString(32)
	if err != nil {
		fmt.Println("Error ustart_go/uses/resendVerificationEmail line 16: Error generating token")
		fmt.Println(err)
		return
	}

	userID, err1 := getUser.UserIDByEmail(eclient, email)
	if err1 != nil {
		fmt.Println("Error ustart_go/uses/resendVerificationEmail line 24: Error finding user by email")
		fmt.Println(err1)
		return
	}

	//Todo: Get user.type by email instead of by id
	user, err2 := getUser.UserByID(eclient, userID)
	if err2 != nil {
		fmt.Println("Error ustart_go/uses/resendVerificationEmail line 31: Error getting user by ID")
		fmt.Println(err1)
		return
	}

	if !user.FirstLogin {
		err3 := postUser.UpdateUser(eclient, userID, "AuthenticationCode", token)
		if err2 != nil {
			fmt.Println("Error ustart_go/uses/resendVerificationEmail line 40: Error resetting Authentication Code")
			fmt.Println(err3)
			return
		}

		from := "ustarttestemail@gmail.com"
		pass := "Ust@rt20!8~~"
		body := "http://ustart.today:5002/Activation/?email=" + email + "&verifCode=" + token
		msg := "From: " + from + "\n" + "To: " + email + "\n" + "Subject: UStart Verification Code\n\n" + body

		err4 := smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{email}, []byte(msg))
		if err4 != nil {
			log.Printf("smtp error: %s", err4)
			return
		}
		// SendEmail(email, token)
		//Todo: Fix email formatting
	}
	return
}
