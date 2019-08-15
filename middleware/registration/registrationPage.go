package registration

import (
	"errors"

	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	getBadge "github.com/sea350/ustart_go/get/badge"
	getGC "github.com/sea350/ustart_go/get/guestCode"
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
	p := bluemonday.UGCPolicy()

	// if len(r.FormValue("inputEmail")) == 0 {
	// 	return
	// }

	//proper email
	if !uses.ValidEmail(p.Sanitize(r.FormValue("inputEmail"))) {

		client.Logger.Println("Email: " + p.Sanitize(r.FormValue("inputEmail")) + " | " + "Invalid email submitted")
		cs := client.ClientSide{ErrorOutput: errors.New("Invalid email submitted"), ErrorStatus: true}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-reg-nil", cs)
		return

	}

	//proper username
	if !uses.ValidUsername(r.FormValue("username")) {

		client.Logger.Println("DocID: " + p.Sanitize(r.FormValue("inputEmail")) + " | " + "Invalid username submitted")
		cs := client.ClientSide{ErrorOutput: errors.New("Invalid username submitted"), ErrorStatus: true}
		client.RenderTemplate(w, r, "templateNoUser2", cs)
		client.RenderTemplate(w, r, "new-reg-nil", cs)
		return

	}

	//	u.FirstName = r.FormValue("firstName")
	fname := p.Sanitize(r.FormValue("firstName"))
	lname := p.Sanitize(r.FormValue("lastName"))
	email := strings.ToLower(p.Sanitize(r.FormValue("inputEmail")))

	username := p.Sanitize(r.FormValue("username"))

	password := r.FormValue("inputPassword")
	passwordb := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordb, bcrypt.DefaultCost)
	school := p.Sanitize(r.FormValue("universityName"))
	var major []string
	major = append(major, p.Sanitize(r.FormValue("majors")))

	var bday time.Time
	if len(r.FormValue("dob")) != 0 {
		year, _ := strconv.Atoi(r.FormValue("dob")[0:4])
		month, _ := strconv.Atoi(r.FormValue("dob")[5:7])
		day, _ := strconv.Atoi(r.FormValue("dob")[8:10])
		bday = time.Date(year, time.Month(month), day, 1, 1, 1, 1, time.UTC)
		//proper birth date
		if !uses.ValidDate(r.FormValue("dob")) {

			client.Logger.Println("DocID: " + p.Sanitize(r.FormValue("inputEmail")) + " | " + "Invalid birthdate submitted")
			cs := client.ClientSide{ErrorOutput: errors.New("Invalid birth date submitted"), ErrorStatus: true}
			client.RenderTemplate(w, r, "templateNoUser2", cs)
			client.RenderTemplate(w, r, "new-reg-nil", cs)
			return

		}
	}
	// if bday == time.Now() {
	// 	log.Println(bday)
	// }
	// country := r.FormValue("country")
	// state := r.FormValue("state")
	// city := p.Sanitize(r.FormValue("city"))
	// zip := p.Sanitize(r.FormValue("zip"))
	currYear := r.FormValue("year")

	//attempting to catch client IP
	var clientIP string
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {

		client.Logger.Printf("Userip: %q is not IP:port\n", r.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {

		client.Logger.Printf("UserIp: %q is not IP:port\n", r.RemoteAddr)
	} else {
		clientIP = userIP.String()
	}

	ref := r.FormValue("ref")

	//check if URL is valid (url code is legitimate)
	isValid, err := uses.ValidGuestCode(client.Eclient, ref)

	if len(ref) != 0 && err != nil {
		client.Logger.Println("Email: "+email+" | Badge code validation error: ", err)
	}

	// client.Logger.Println(ref)
	// client.Logger.Println(isValid)

	if len(ref) != 0 && isValid {
		err2 := uses.BadgeSignUpBasic(client.Eclient, username, email, hashedPassword, fname, lname, school, major, bday, currYear, clientIP, ref)

		if err2 != nil {

			client.Logger.Println("Email: "+email+" | err at signup: ", err2)
			cs := client.ClientSide{ErrorOutput: err2, ErrorStatus: true}
			client.RenderTemplate(w, r, "templateNoUser2", cs)
			client.RenderTemplate(w, r, "new-reg-nil", cs)

		} else {
			http.Redirect(w, r, "/registrationcomplete/", http.StatusFound)
		}

	} else {

		err3 := uses.SignUpBasic(client.Eclient, username, email, hashedPassword, fname, lname, school, major, bday, currYear, clientIP) // country, state, city, zip,

		if err3 != nil {

			client.Logger.Println("Email: "+email+" | err at signup: ", err3)
			cs := client.ClientSide{ErrorOutput: err3, ErrorStatus: true}
			client.RenderTemplate(w, r, "templateNoUser2", cs)
			client.RenderTemplate(w, r, "new-reg-nil", cs)

		} else {
			http.Redirect(w, r, "/registrationcomplete/", http.StatusFound)
		}
	}

}

//Signup ... tests for an existing docId in sesson, if no id then render signup, if there is id redirect to profile
func Signup(w http.ResponseWriter, r *http.Request) {
	client.Store.MaxAge(8640 * 7)
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	cs := client.ClientSide{}
	if test1 != nil {
		http.Redirect(w, r, "/profile/"+test1.(string), http.StatusFound)
		return
	}
	//r.ParseForm()

	// client.Logger.Println(r.URL.Path)
	ref := r.FormValue("ref")

	isValid, err := uses.ValidGuestCode(client.Eclient, ref)
	if len(ref) != 0 && err != nil {
		client.Logger.Println("Reference: "+ref+" | err at signup: ", err)
		// cs.ErrorStatus = true
		// cs.ErrorOutput = errors.New("Invalid reference code")
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	if len(ref) != 0 && isValid {

		gcObject, err := getGC.GuestCodeByID(client.Eclient, ref)
		if err != nil {
			client.Logger.Println("Reference: "+ref+" | err at signup: ", err)
			http.Redirect(w, r, "/404/", http.StatusFound)
			return
		}

		badge, err := getBadge.BadgeByID(client.Eclient, gcObject.Description)
		if err != nil {
			client.Logger.Println("Reference: "+ref+" | err at signup: ", err)
			http.Redirect(w, r, "/404/", http.StatusFound)
			return
		}

		if len(badge.ID) == 0 {
			client.Logger.Println("Reference: "+ref+" | err at signup: ", err)
			http.Redirect(w, r, "/404/", http.StatusFound)
			return
		}
	}
	session.Save(r, w)
	client.RenderTemplate(w, r, "templateNoUser2", cs)
	client.RenderTemplate(w, r, "new-reg-nil", cs)
}
