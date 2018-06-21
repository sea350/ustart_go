package registration

import (
	"errors"
	"fmt"
	"net/http"

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

	user, err1 := get.UserByEmail(client.Eclient, email)
	if err1 != nil {
		fmt.Println("err: middleware/registration/emailVerification line 22")
		fmt.Println(err1)
		cs.ErrorStatus = true
		cs.ErrorOutput = err1
		return
	}

	if emailToken == user.AuthenticationCode {
		userID, err := get.UserIDByEmail(client.Eclient, email)
		if err != nil {
			fmt.Println("err: middleware/registration/emailVerification line 32")
			fmt.Println(err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
			return
		}

		err2 := post.UpdateUser(client.Eclient, userID, "FirstLogin", true)
		if err2 != nil {
			fmt.Println("err: middleware/registration/emailVerification line 41")
			fmt.Println(err2)
			cs.ErrorStatus = true
			cs.ErrorOutput = err2
			return
		}
		err3 := post.UpdateUser(client.Eclient, userID, "AuthenticationCode", nil)
		if err3 != nil {
			fmt.Println("err: middleware/registration/emailVerification line 49")
			fmt.Println(err3)
			cs.ErrorStatus = true
			cs.ErrorOutput = errors.New("err: middleware/registration/emailVerification line 49")
		}
	} else {
		fmt.Println("Verification code is wrong")
		cs.ErrorStatus = true
		cs.ErrorOutput = errors.New("Verification code is wrong")
	}
	return
}
