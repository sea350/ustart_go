package registration

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

func EmailVerification(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	emailToken := r.FormValue("verifCode")

	userID, err := get.UserIDByEmail(client.Eclient, email)
	if err != nil {
		fmt.Println("err: middleware/registration/emailVerification line 16")
		fmt.Println(err)
		return
	}

	user, err1 := get.UserByEmail(client.Eclient, email)
	if err1 != nil {
		fmt.Println("err: middleware/registration/emailVerification line 22")
		fmt.Println(err1)
		return
	}

	if emailToken == user.AuthenticationCode {
		err2 := post.UpdateUser(client.Eclient, userID, "FirstLogin", true)
		if err2 != nil {
			fmt.Println("err: middleware/registration/emailVerification line 29")
			fmt.Println(err2)
			return
		}
		err3 := post.UpdateUser(client.Eclient, userID, "AuthenticationCode", nil)
		if err3 != nil {
			fmt.Println("err: middleware/registration/emailVerification line 34")
			fmt.Println(err3)
			return
		}
	} else {
		fmt.Println("Verification code is wrong")
		return
	}

}
