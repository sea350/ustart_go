package login

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	client "github.com/sea350/ustart_go/middleware/clientstruct"
	uses "github.com/sea350/ustart_go/uses"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // represents form data from html
	session, _ := store.Get(r, "session_please")
	// check if docid exists within the session note: there is inconsistency with checking docid/username.
	test1, _ := session.Values["Username"]
	if test1 != nil {
		http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	}
	email := r.FormValue("email")
	email = strings.ToLower(email) // we only store lowercase emails in the db
	//	var password []byte
	password := r.FormValue("password")
	//	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	passwordb := []byte(password)

	successful, sessionInfo, err2 := uses.Login(eclient, email, passwordb)

	// doc ID can be retrieved here!
	//cs := &ClientSide{}

	if err2 != nil {
		fmt.Println(err2)
		fmt.Println("This is an error, LoginFile.go: 40")
	}

	if successful {
		session.Values["DocID"] = sessionInfo.DocID
		session.Values["FirstName"] = sessionInfo.FirstName
		session.Values["LastName"] = sessionInfo.LastName
		session.Values["Email"] = sessionInfo.Email
		session.Values["Username"] = sessionInfo.Username
		expiration := time.Now().Add((30) * time.Hour)
		fmt.Println(sessionInfo)
		cookie := http.Cookie{Name: session.Values["DocID"].(string), Value: "user", Expires: expiration, Path: "/"}
		http.SetCookie(w, &cookie)
		session.Save(r, w)
		http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	}

	if !successful {
		cs := client.ClientSide{ErrorLogin: true}
		fmt.Println(successful)
		fmt.Println("This is an error, LoginFile.go: 55")
		client.RenderTemplate(w, "templateNoUser2", cs)
		client.RenderTemplate(w, "loginerror-nil", cs)

	}
}

// func LoggedIn (w http.ResponseWriter, r *http.Request){
// 	session, _ := store.Get(r, "session_please")
// 	//	fmt.Println(session.Values["FirstName"].(string))
// 	cs := ClientSide{FirstName:session.Values["FirstName"].(string)}
// 	session.Save(r, w)
// 	renderTemplate(w,"template2-nil",cs)
// 	renderTemplate(w,"profile-nil",cs)
// }

/* This isn't used anymore I think*/
func LoginError(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		http.Redirect(w, r, "/profile/"+session.Values["DocID"].(string), http.StatusFound)
	}
	email := r.FormValue("email")
	//	var password []byte
	password := r.FormValue("password")
	fmt.Println(password)
	//	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	passwordb := []byte(password)
	successful, sessionInfo, err2 := uses.Login(eclient, email, passwordb)
	if err2 != nil {
		fmt.Println(err2)

	}

	if successful == true {
		fmt.Println("login successful")
		session.Values["DocID"] = sessionInfo.DocID
		session.Values["FirstName"] = sessionInfo.FirstName
		session.Values["LastName"] = sessionInfo.LastName
		session.Values["Email"] = sessionInfo.Email
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
		//	cs := ClientSide{ErrorLogin: errorL}
		fmt.Println("errorL is ")
		fmt.Print(errorL)
		http.Redirect(w, r, "/loginerror-nil/", http.StatusFound)

	}

}
