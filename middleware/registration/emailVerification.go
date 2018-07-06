package registration

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

func EmailVerification(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	emailToken := r.FormValue("verifCode")

	var cs client.ClientSide

	defer client.RenderSidebar(w, r, "template2-nil")
	defer client.RenderTemplate(w, r, "reg-got-verified", cs)

	userID, err := get.UserIDByEmail(client.Eclient, email)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		cs.ErrorStatus = true
		cs.ErrorOutput = err
		return
	}

	user, err := get.UserByEmail(client.Eclient, email)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		cs.ErrorStatus = true
		cs.ErrorOutput = err
		return
	}

	if emailToken == user.AuthenticationCode {
		err = post.UpdateUser(client.Eclient, userID, "FirstLogin", true)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
			return
		}
		err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", nil)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			cs.ErrorStatus = true
			cs.ErrorOutput = errors.New("err: middleware/registration/emailVerification line 34")
		}
	} else {
		fmt.Println("Verification code is wrong")
		cs.ErrorStatus = true
		cs.ErrorOutput = errors.New("Verification code is wrong")
	}

}
