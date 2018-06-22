package uses

import (
	"fmt"
	"log"
	"net/smtp"

	elastic "gopkg.in/olivere/elastic.v5"
)

//SendPasswordResetEmail ... Sends password reset token link to user
//Requires the user's email address
//Returns if the email failed to send
func SendPasswordResetEmail(eclient *elastic.Client, email string) {
	token, err := GenerateRandomString(32)
	if err != nil {
		fmt.Println("Error ustart_go/uses/resendVerificationEmail line 16: Error generating token")
		fmt.Println(err)
		return
	}

	from := "ustarttestemail@gmail.com"
	pass := "Ust@rt20!8~~"
	body := "http://ustart.today:5002/ResetPassword/?email=" + email + "&verifCode=" + token
	msg := "From: " + from + "\n" + "To: " + email + "\n" + "Subject: UStart Password Reset\n\n" + body

	err4 := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{email}, []byte(msg))
	if err4 != nil {
		log.Printf("smtp error: %s", err4)
		return
	}
	SendEmail(email, token)
}
