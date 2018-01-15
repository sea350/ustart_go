package client

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//Eclient ... Reference to the ElasticSearch
var Eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

//Store ...
var Store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

//ClientSide ... This struct represents a user state after he/she has logged in. Some fields may no longer be needed
//or are unnecessary.
type ClientSide struct {
	DOCID           string
	FirstName       string
	LastName        string
	Username        string
	ErrorOutput     error
	ErrorStatus     bool
	UserInfo        types.User
	Class           string
	Birthday        string
	ImageCode       string
	Description     string
	Followers       int
	Following       int
	Page            string
	FollowingStatus string
	Wall            []types.JournalEntry
}

/* The following line is how HTML is loaded by our application. Note we need the relative link from the location of GoStart2. */
var templates = template.Must(template.ParseFiles("/home/rr2396/www/ustart.tech/followerlist-nil.html",
	"/home/rr2396/www/ustart.tech/emTee.html", "/home/rr2396/www/ustart.tech/wallttt.html",
	"/home/rr2396/www/ustart.tech/wallload-nil.html", "/home/rr2396/www/ustart.tech/testimage.html",
	"/home/rr2396/www/ustart.tech/ajax-nil.html", "/home/rr2396/www/ustart.tech/Membership-Nil.html",
	"/home/rr2396/www/ustart.tech/settings-Nil.html", "/home/rr2396/www/ustart.tech/inbox-Nil.html",
	"/home/rr2396/www/ustart.tech/createProject-Nil.html", "/home/rr2396/www/ustart.tech/manageprojects-Nil.html",
	"/home/rr2396/www/ustart.tech/projectsF.html", "/home/rr2396/www/ustart.tech/new-reg-nil.html",
	"/home/rr2396/www/ustart.tech/loginerror-nil.html", "/home/rr2396/www/ustart.tech/test.html",
	"/home/rr2396/www/ustart.tech/payment-nil.html", "/home/rr2396/www/ustart.tech/templateNoUser2.html",
	"/home/rr2396/www/ustart.tech/profile-nil.html", "/home/rr2396/www/ustart.tech/template2-nil.html",
	"/home/rr2396/www/ustart.tech/template-footer-nil.html", "/home/rr2396/www/ustart.tech/nil-index2.html",
	"/home/rr2396/www/ustart.tech/regcomplete-nil.html"))

//RenderTemplate ... This function does the actual rendering of HTML pages. Note it takes in a struct (type ClientSide).
//You will need to continually send data to the pages and this is accomplished via the struct.
func RenderTemplate(w http.ResponseWriter, tmpl string, cs ClientSide) {
	err := templates.ExecuteTemplate(w, tmpl+".html", cs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
