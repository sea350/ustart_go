package login

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	uses "github.com/sea350/ustart_go/uses"
)

//Login ... logs you in duh
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // represents form data from html
	session, _ := client.Store.Get(r, "session_please")
	// check if docid exists within the session note: there is inconsistency with checking docid/username.
	test1, _ := session.Values["Username"]
	if test1 != nil {
		http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	}

	//attempting to catch client IP
	var clientIP string
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Printf("userip: %q is not IP:port\n", r.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Printf("userip: %q is not IP:port\n", r.RemoteAddr)
	} else {
		clientIP = userIP.String()
	}

	email := r.FormValue("email")
	email = strings.ToLower(email) // we only client.Store lowercase emails in the db
	//	var password []byte
	password := r.FormValue("password")
	//	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	passwordb := []byte(password)

	successful, sessionInfo, err2 := uses.Login(client.Eclient, email, passwordb, clientIP)

	// doc ID can be retrieved here!
	//cs := &ClientSide{}

	if err2 != nil {
		fmt.Println(err2)
		fmt.Println("This is an error, LoginFile.go: 40")
	}

	if successful {
		session.Values["DocID"] = sessionInfo.DocID
		fmt.Println("user logged in: " + sessionInfo.DocID)
		session.Values["FirstName"] = sessionInfo.FirstName
		session.Values["LastName"] = sessionInfo.LastName
		session.Values["Email"] = sessionInfo.Email
		session.Values["Username"] = sessionInfo.Username
		//session.Values["Avatar"] = sessionInfo.Avatar
		expiration := time.Now().Add((30) * time.Hour)
		cookie := http.Cookie{Name: session.Values["DocID"].(string), Value: "user", Expires: expiration, Path: "/"}
		http.SetCookie(w, &cookie)
		session.Save(r, w)
		http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

		//Deletes authentication code (if any) from user
		userID, err := get.UserIDByEmail(client.Eclient, email)
		if err != nil {
			fmt.Println("Error: /ustart_go/middleware/login/LoginFile/ line 70: Unable to retrieve user")
			fmt.Println(err)
		} else {
			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", nil)
			if err != nil {
				fmt.Println("Error: /ustart_go/middleware/login/LoginFile/ line 76: Unable to remove authentication code")
				fmt.Println(err)
			}
			err = post.UpdateUser(client.Eclient, userID, "AuthenticationCodeTime", nil)
			if err != nil {
				fmt.Println("Error: /ustart_go/middleware/login/LoginFile/ line 81: Unable to remove authentication code time")
				fmt.Println(err)
			}
		}
		return
	}

	if !successful {
		cs := client.ClientSide{ErrorStatus: true}
		fmt.Println(successful)
		fmt.Println("This is an error, LoginFile.go: 55")
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "loginerror-nil", cs)

	}
}

// func LoggedIn (w http.ResponseWriter, r *http.Request){
// 	session, _ := client.Store.Get(r, "session_please")
// 	//	fmt.Println(session.Values["FirstName"].(string))
// 	cs := ClientSide{FirstName:session.Values["FirstName"].(string)}
// 	session.Save(r, w)
// 	renderTemplate(w,"template2-nil",cs)
// 	renderTemplate(w,"profile-nil",cs)
// }

//Error ... This isn't used anymore I think
func Error(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		http.Redirect(w, r, "/profile/"+session.Values["DocID"].(string), http.StatusFound)
	}
	email := r.FormValue("email")
	//	var password []byte
	password := r.FormValue("password")
	fmt.Println("DEBUG: middleware/LoginFile.go line: 80-81")
	fmt.Println(password)
	//	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	passwordb := []byte(password)
	successful, sessionInfo, err2 := uses.Login(client.Eclient, email, passwordb, "0")
	if err2 != nil {
		fmt.Println(err2)

	}

	if successful == true {
		fmt.Println("login successful")
		session.Values["DocID"] = sessionInfo.DocID
		session.Values["FirstName"] = sessionInfo.FirstName
		session.Values["LastName"] = sessionInfo.LastName
		session.Values["Email"] = sessionInfo.Email
		//session.Values["Avatar"] = sessionInfo.Avatar
		expiration := time.Now().Add((30) * time.Hour)
		cookie := http.Cookie{Name: session.Values["DocID"].(string), Value: "user", Expires: expiration, Path: "/"}
		http.SetCookie(w, &cookie)
		session.Save(r, w)
		http.Redirect(w, r, "/profile/"+session.Values["DocID"].(string), http.StatusFound)
	}

	if successful == false {
		fmt.Println("did not login successful")
		var errorL bool
		errorL = true
		fmt.Println("errorL is ")
		fmt.Print(errorL)
		http.Redirect(w, r, "/loginerror-nil/", http.StatusFound)

	}

}
