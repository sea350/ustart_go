package uses

import (
	"log"
	"time"

	"github.com/sea350/ustart_go/globals"

	elastic "github.com/olivere/elastic"
	getUser "github.com/sea350/ustart_go/get/user"
	postUser "github.com/sea350/ustart_go/post/user"
)

//SendVerificationEmail ... resends User Verification Email
//Requires the user's email address
//Returns if the email failed to send
func SendVerificationEmail(eclient *elastic.Client, email string) {

	time.Sleep(30 * time.Second)

	userID, err := getUser.UserIDByEmail(eclient, email)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	//Todo: Get user.type by email instead of by id
	user, err := getUser.UserByID(eclient, userID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	if !user.Verified {
		token, err := GenerateRandomString(32)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}

		err = postUser.UpdateUser(eclient, userID, "AuthenticationCode", token)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}

		err = postUser.UpdateUser(eclient, userID, "AuthenticationCodeTime", time.Now())
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}

		subject := "Your verification link"
		// link := globals.SiteURL + ":" + globals.Port + "/Activation/?email=" + email + "&verifCode=" + token
		link := globals.SiteURL + "/Activation/?email=" + email + "&verifCode=" + token
		r := NewRequest([]string{email}, subject)
		r.Send(globals.HTMLPATH+"email_template.html", map[string]string{"username": user.Username, "link": link,
			"contentjuan":   "We received a request to activate your Ustart Account. We would love to assist you!",
			"contentdos":    "Simply click the button below to verify your account",
			"contenttres":   "VERIFY ACCOUNT",
			"contentquatro": "a new account"})
	}
	return
}
