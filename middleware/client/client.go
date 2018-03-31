package client

import (
	htype "html/template"
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
	ListOfHeads     []types.FloatingHead
	ListOfHeads2    []types.FloatingHead
	Wall            []types.JournalEntry
	Widgets         []types.Widget
	Project         types.ProjectAggregate
}

/* The following line is how HTML is loaded by our application. Note we need the relative link from the location of GoStart2. */
var templates = htype.Must(htype.ParseFiles("/ustart/ustart_front/followerlist-nil.html",
	"/ustart/ustart_front/emTee.html", "/ustart/ustart_front/wallttt.html",
	"/ustart/ustart_front/wallload-nil.html", "/ustart/ustart_front/testimage.html",
	"/ustart/ustart_front/ajax-nil.html", "/ustart/ustart_front/Membership-Nil.html",
	"/ustart/ustart_front/settings-Nil.html", "/ustart/ustart_front/inbox-Nil.html",
	"/ustart/ustart_front/createProject-Nil.html", "/ustart/ustart_front/manageprojects-Nil.html",
	"/ustart/ustart_front/projectsF.html", "/ustart/ustart_front/new-reg-nil.html",
	"/ustart/ustart_front/loginerror-nil.html", "/ustart/ustart_front/test.html",
	"/ustart/ustart_front/payment-nil.html", "/ustart/ustart_front/templateNoUser2.html",
	"/ustart/ustart_front/profile-nil.html", "/ustart/ustart_front/template2-nil.html",
	"/ustart/ustart_front/template-footer-nil.html", "/ustart/ustart_front/nil-index2.html",
	"/ustart/ustart_front/regcomplete-nil.html", "/ustart/ustart_front/project_settings_F.html",
	"/ustart/ustart_front/leftnav-nil.html", "/ustart/ustart_front/ManageProjectMembersF.html",
	"/ustart/ustart_front/followerlist-nil.html", "/ustart/ustart_front/404.html"))

//RenderTemplate ... This function does the actual rendering of HTML pages. Note it takes in a struct (type ClientSide).
//You will need to continually send data to the pages and this is accomplished via the struct.
func RenderTemplate(w http.ResponseWriter, tmpl string, cs ClientSide) {
	err := templates.ExecuteTemplate(w, tmpl+".html", cs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
