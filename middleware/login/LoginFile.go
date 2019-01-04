package login

import (
	"net"
	"net/http"
	"strings"
	"time"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
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
		client.Logger.Printf("userip: %q is not IP:port\n", r.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		client.Logger.Printf("userip: %q is not IP:port\n", r.RemoteAddr)
	} else {
		clientIP = userIP.String()
	}

	email := r.FormValue("email")
	email = strings.ToLower(email) // we only client.Store lowercase emails in the db
	password := r.FormValue("password")
	passwordb := []byte(password)

	//Check if user is verified
	user, _ := get.UserByEmail(client.Eclient, email)

	successful, sessionInfo, err := uses.Login(client.Eclient, email, passwordb, clientIP)

	if !successful || err != nil {

		client.Logger.Println("Email: "+email+" | err: ", err)
		client.RenderTemplate(w, r, "templateNoUser2", client.ClientSide{ErrorStatus: true, ErrorOutput: err})
		client.RenderTemplate(w, r, "loginerror-nil", client.ClientSide{ErrorStatus: true, ErrorOutput: err})
		return
	}

	if !user.Verified {
		session.Values["Email"] = user.Email
		session.Save(r, w)
		http.Redirect(w, r, "/unverified/", http.StatusFound)

		client.Logger.Println("DocID: " + sessionInfo.DocID + " | " + "User not verified")
		return
	}

	if successful {
		session.Values["DocID"] = sessionInfo.DocID
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
	}

}

// func LoggedIn (w http.ResponseWriter, r *http.Request){
// 	session, _ := client.Store.Get(r, "session_please")
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

	//	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	passwordb := []byte(password)

	//attempting to catch client IP
	var clientIP string
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		client.Logger.Printf("userip: %q | err: %s \n", r.RemoteAddr, err)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		client.Logger.Printf("userip: %q | err: %s \n", r.RemoteAddr, err)
	} else {
		clientIP = userIP.String()
	}

	successful, sessionInfo, err2 := uses.Login(client.Eclient, email, passwordb, clientIP)
	if err2 != nil {
		client.Logger.Printf("userip: %q | err: %s \n", r.RemoteAddr, err2)
	}

	if successful == true {
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
		http.Redirect(w, r, "/loginerror-nil/", http.StatusFound)
	}

}
