package registration

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
	"golang.org/x/crypto/bcrypt"
)

//Complete ...
func Complete(w http.ResponseWriter, r *http.Request) {
	cs := client.ClientSide{}
	client.RenderTemplate(w, r, "templateNoUser2", cs)
	client.RenderTemplate(w, r, "regcomplete-nil", cs)
}

//RegisterType ...
func RegisterType(w http.ResponseWriter, r *http.Request) {
	cs := client.ClientSide{}
	client.RenderTemplate(w, r, "templateNoUser2", cs)
	client.RenderTemplate(w, r, "Membership-Nil", cs)
}

//Registration ...
func Registration(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session, _ := client.Store.Get(r, "session_please")
	// check DOCID instead
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		// REGISTRATION SHOULD NOT LOG YOU IN
		http.Redirect(w, r, "/profile/", http.StatusFound)
		return
	}

	if !uses.ValidEmail(r.FormValue("inputEmail")) {
		fmt.Println("This is an error: registrationPage.go, 45")
		fmt.Println("Invalid email submitted")
		cs := client.ClientSide{ErrorOutput: errors.New("Invalid email submitted"), ErrorStatus: true}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-reg-nil", cs)
		return

	}

	if !uses.ValidUsername(r.FormValue("username")) {
		fmt.Println("This is an error: registrationPage.go, 43")
		fmt.Println("Invalid username submitted")
		cs := client.ClientSide{ErrorOutput: errors.New("Invalid username submitted"), ErrorStatus: true}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-reg-nil", cs)
		return

	}

	//	u.FirstName = r.FormValue("firstName")
	fname := r.FormValue("firstName")
	lname := r.FormValue("lastName")
	email := r.FormValue("inputEmail")
	email = strings.ToLower(email)

	username := r.FormValue("username")

	password := r.FormValue("inputPassword")
	passwordb := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordb, bcrypt.DefaultCost)
	school := r.FormValue("universityName")
	var major []string
	major = append(major, r.FormValue("majors"))
	fmt.Println(r.FormValue("dob"))
	bday := time.Now() //r.FormValue("dob")
	year, _ := strconv.Atoi(r.FormValue("dob")[0:4])
	month, _ := strconv.Atoi(r.FormValue("dob")[5:7])
	day, _ := strconv.Atoi(r.FormValue("dob")[8:10])
	bday = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	fmt.Println(bday.Date())
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	currYear := r.FormValue("year")

	err2 := uses.SignUpBasic(client.Eclient, username, email, hashedPassword, fname, lname, country, state, city, zip, school, major, bday, currYear)

	if err2 != nil {
		fmt.Println("This is an error: registrationPage.go, 65")
		fmt.Println(err2)
		cs := client.ClientSide{ErrorOutput: err2, ErrorStatus: true}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-reg-nil", cs)

	}

	if err2 == nil {
		http.Redirect(w, r, "/registrationcomplete/", http.StatusFound)
		return
	}

}

//Signup ... tests for an existing docId in sesson, if no id then render signup, if there is id redirect to profile
func Signup(w http.ResponseWriter, r *http.Request) {
	client.Store.MaxAge(8640 * 7)
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]

	if test1 != nil {
		fmt.Println(test1)
		fmt.Println("this is debug code: registrationPage.go 89")
		http.Redirect(w, r, "/profile/"+test1.(string), http.StatusFound)
		return
	}

	session.Save(r, w)
	cs := client.ClientSide{ErrorStatus: false}
	client.RenderTemplate(w, r, "templateNoUser2", cs)
	client.RenderTemplate(w, r, "new-reg-nil", cs)
}
