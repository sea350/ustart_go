package registration

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//EmailVerification ... Takes in W and R
//Takes in the form data (email and auth-token) from user and checks if it is correct
//Returns error if false
func EmailVerification(w http.ResponseWriter, r *http.Request) {
	p := bluemonday.UGCPolicy()
	email := p.Sanitize(r.FormValue("email"))
	emailToken := p.Sanitize(r.FormValue("verifCode"))

	var cs client.ClientSide

	defer client.RenderSidebar(w, r, "templateNoUser2")
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
		err = post.UpdateUser(client.Eclient, userID, "Verified", true)
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
