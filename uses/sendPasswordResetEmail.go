package uses

import (
	"fmt"
	"log"
	"net/smtp"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//SendPasswordResetEmail ... Sends password reset token link to user and saves token to their AuthenticationCode
//Requires a valid user email address
//Returns if the email failed to send
func SendPasswordResetEmail(eclient *elastic.Client, email string) {
	userID, err := get.UserIDByEmail(client.Eclient, email)
	if err != nil {
		fmt.Println("Error ustart_go/uses/resendVerificationEmail line 16: Unable to retreive userID by email")
		fmt.Println(err)
		return
	}

	token, err := GenerateRandomString(32)
	if err != nil {
		fmt.Println("Error ustart_go/uses/resendVerificationEmail line 16: Error generating token")
		fmt.Println(err)
		return
	}

	err = post.UpdateUser(eclient, userID, "AuthenticationCode", token)
	if err != nil {
		fmt.Println("err: Unable to retreive userID by email")
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
}
