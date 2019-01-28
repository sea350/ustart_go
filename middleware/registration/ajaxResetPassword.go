package registration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
	"golang.org/x/crypto/bcrypt"
)

//AjaxResetPassword ... Reset's user's password
//Requires the user's email address
//Returns if the email failed to send
func AjaxResetPassword(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()

	email := r.FormValue("email")
	token := r.FormValue("verifCode")
	pass := r.FormValue("password")

	res := make(map[string]string)

	if email == `` || token == `` || pass == `` {
		res["error"] = "Critical data not passed in"
		data, err := json.Marshal(res)
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}

		fmt.Fprintln(w, string(data))
		return
	}

	user, err := get.UserByEmail(client.Eclient, email)
	if err != nil {
		client.Logger.Println("Email: "+email+" | err: ", err)
		res["error"] = err.Error()
		data, err := json.Marshal(res)
		if err != nil {
			client.Logger.Println("Email: "+email+" | err: ", err)
		}

		fmt.Fprintln(w, string(data))
		return
	}

	// If the time since authentication code was input is less than 1 hour
	if time.Since(user.AuthenticationCodeTime) < (time.Second*3600) && token == user.AuthenticationCode {
		newHashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
		if err != nil {
			client.Logger.Println("Email: "+email+" | err: ", err)
			res["error"] = err.Error()
			data, err := json.Marshal(res)
			if err != nil {
				client.Logger.Println("Email: "+email+" | err: ", err)
			}

			fmt.Fprintln(w, string(data))
			return
		}

		userID, err := get.UserIDByEmail(client.Eclient, email)
		if err != nil {
			client.Logger.Println("Email: "+email+" | err: ", err)
			res["error"] = err.Error()
			data, err := json.Marshal(res)
			if err != nil {
				client.Logger.Println("Email: "+email+" | err: ", err)
			}

			fmt.Fprintln(w, string(data))
			return
		}

		user.Password = newHashedPass
		// err = post.UpdateUser(client.Eclient, userID, "Password", newHashedPass)
		// if err != nil {

		// 	client.Logger.Println("DocID: "+userID+" | err: ", err)
		// 	cs.ErrorOutput = errors.New("A problem has occurred")
		// 	cs.ErrorStatus = true
		// 	client.RenderSidebar(w, r, "templateNoUser2")
		// 	client.RenderTemplate(w, r, "reset-new-pass", cs)
		// 	return
		// }

		user.AuthenticationCode = ""

		// err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", nil)
		// if err != nil {

		// 	client.Logger.Println("DocID: "+userID+" | err: ", err)
		// }
		var newTime time.Time
		user.AuthenticationCodeTime = newTime

		// err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", nil)
		// if err != nil {

		// 	client.Logger.Println("DocID: "+userID+" | err: ", err)
		// }

		user.LoginWarnings = make(map[string]types.LoginWarning)

		// err = post.UpdateUser(client.Eclient, userID, "LoginWarnings", make(map[string]types.LoginWarning))
		err = post.ReindexUser(client.Eclient, userID, user)
		if err != nil {

			client.Logger.Println("DocID: "+userID+" | err: ", err)
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}

	client.Logger.Println("Email: " + email + " | err: " + "Authentication token invalid")
	res["error"] = "Authentication token invalid"
	data, err := json.Marshal(res)
	if err != nil {
		client.Logger.Println("Email: "+email+" | err: ", err)
	}

	fmt.Fprintln(w, string(data))

}
