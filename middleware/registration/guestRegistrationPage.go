package registration

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	client "github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/uses"
	bcrypt "golang.org/x/crypto/bcrypt"
)

//GuestRegistration ... Separate registration page for guests (non-NYU users)
func GuestRegistration(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		http.Redirect(w, r, "/profile/", http.StatusFound)
		return
	}
	p := bluemonday.UGCPolicy()

	fname := p.Sanitize(r.FormValue("firstName"))
	lname := p.Sanitize(r.FormValue("lastName"))
	email := strings.ToLower(p.Sanitize(r.FormValue("inputEmail")))
	username := p.Sanitize(r.FormValue("username"))
	password := r.FormValue("inputPassword")
	guestCode := p.Sanitize(r.FormValue("guestCode"))
	passwordb := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordb, bcrypt.DefaultCost)
	school := p.Sanitize(r.FormValue("universityName"))
	var major []string
	major = append(major, p.Sanitize(r.FormValue("majors")))
	year, _ := strconv.Atoi(r.FormValue("dob")[0:4])
	month, _ := strconv.Atoi(r.FormValue("dob")[5:7])
	day, _ := strconv.Atoi(r.FormValue("dob")[8:10])
	bday := time.Date(year, time.Month(month), day, 1, 1, 1, 1, time.UTC)
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := p.Sanitize(r.FormValue("city"))
	zip := p.Sanitize(r.FormValue("zip"))
	currYear := r.FormValue("year")

	//if the form is blank just render w/o returning errors
	if email == `` && username == `` {
		cs := client.ClientSide{}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-guest-reg", cs)
		return
	}

	//proper email
	if !uses.ValidGuestEmail(emial) {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Invalid email submitted")
		cs := client.ClientSide{ErrorOutput: errors.New("Invalid email submitted"), ErrorStatus: true}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-guest-reg", cs)
		return
	}

	//proper username
	if !uses.ValidUsername(username) {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Invalid username submitted")
		cs := client.ClientSide{ErrorOutput: errors.New("Invalid username submitted"), ErrorStatus: true}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-guest-reg", cs)
		return

	}

	//proper birth date
	if !uses.ValidDate(r.FormValue("dob")) {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Invalid date of birth submitted")
		cs := client.ClientSide{ErrorOutput: errors.New("Invalid birth date submitted"), ErrorStatus: true}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-guest-reg", cs)
		return
	}

	err2 := uses.GuestSignUpBasic(client.Eclient, username, email, hashedPassword, fname, lname, country, state, city, zip, school, major, bday, currYear, guestCode)
	if err2 != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err2)
		cs := client.ClientSide{ErrorOutput: err2, ErrorStatus: true}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-guest-reg", cs)

	}

	if err2 == nil {
		http.Redirect(w, r, "/registrationcomplete/", http.StatusFound)
		return
	}

}

//GuestSignup ... tests for an existing docId in sesson, if no id then render signup, if there is id redirect to profile
func GuestSignup(w http.ResponseWriter, r *http.Request) {
	client.Store.MaxAge(8640 * 7)
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]

	if test1 != nil {
		fmt.Println(test1)
		fmt.Println("this is debug code: guestRegistrationPage.go 105")
		http.Redirect(w, r, "/profile/"+test1.(string), http.StatusFound)
		return
	}

	session.Save(r, w)
	cs := client.ClientSide{ErrorStatus: false}
	client.RenderTemplate(w, r, "templateNoUser2", cs)
	client.RenderTemplate(w, r, "new-guest-reg", cs)
}