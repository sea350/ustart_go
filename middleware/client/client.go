package client

import (
	htype "html/template"
	"net/http"

	sessions "github.com/gorilla/sessions"
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
	Avatar          string
	ErrorOutput     error
	ErrorStatus     bool
	UserInfo        types.User
	Class           string
	Birthday        string
	ImageCode       string
	Description     string
	Followers       int
	Following       int
	Page            string //DocID of current page
	FollowingStatus string
	ListOfHeads     []types.FloatingHead
	ListOfHeads2    []types.FloatingHead
	Wall            []types.JournalEntry
	Widgets         []types.Widget
	Project         types.ProjectAggregate
	Event           types.EventAggregate
	Messages        []types.Message
}

/* The following line is how HTML is loaded by our application. Note we need the relative link from the location of GoStart2. */
var templates = htype.Must(htype.ParseFiles("/ustart/ustart_front/followerlist-nil.html", "/ustart/ustart_front/Membership-Nil.html",
	"/ustart/ustart_front/settings-Nil.html", "/ustart/ustart_front/inbox-Nil.html", "/ustart/ustart_front/nil-index2.html",
	"/ustart/ustart_front/createProject-Nil.html",	"/ustart/ustart_front/manageprojects-Nil.html",
	"/ustart/ustart_front/projectsF.html",	"/ustart/ustart_front/new-reg-nil.html",
	"/ustart/ustart_front/loginerror-nil.html", "/ustart/ustart_front/templateNoUser2.html",
	"/ustart/ustart_front/profile-nil.html", "/ustart/ustart_front/template2-nil.html",
	"/ustart/ustart_front/template-footer-nil.html", "/ustart/ustart_front/regcomplete-nil.html",
	"/ustart/ustart_front/project_settings_F.html",	"/ustart/ustart_front/reset-forgot-pw.html", "/ustart/ustart_front/leftnav-nil.html",
	"/ustart/ustart_front/ManageProjectMembersF.html",	"/ustart/ustart_front/followerlist-nil.html", "/ustart/ustart_front/404.html",
	"/ustart/ustart_front/search-nil.html", "/ustart/ustart_front/reg-got-verified.html",
	"/ustart/ustart_front/reset-new-pass.html", "/ustart/ustart_front/cuzsteventoldmeto.html",
	"/ustart/ustart_front/events.html", "/ustart/ustart_front/eventStart.html", "/ustart/ustart_front/eventManager.html",
	"/ustart/ustart_front/chat.html"))

//RenderTemplate ... This function does the actual rendering of HTML pages. Note it takes in a struct (type ClientSide).
//You will need to continually send data to the pages and this is accomplished via the struct.
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, cs ClientSide) {
	session, _ := Store.Get(r, "session_please")
	if session.Values["FirstName"] != nil {
		cs.FirstName = session.Values["FirstName"].(string)
	}
	if session.Values["LastName"] != nil {
		cs.LastName = session.Values["LastName"].(string)
	}
	if session.Values["Username"] != nil {
		cs.Username = session.Values["Username"].(string)
	}
	if session.Values["DocID"] != nil {
		cs.DOCID = session.Values["DocID"].(string)
	}
	if session.Values["Avatar"] != nil {
		cs.Avatar = session.Values["Avatar"].(string)
	}
	err := templates.ExecuteTemplate(w, tmpl+".html", cs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//RenderSidebar ... used to render anything only needs session information, like sidebar or banner
func RenderSidebar(w http.ResponseWriter, r *http.Request, tmpl string) {
	session, _ := Store.Get(r, "session_please")
	var cs ClientSide
	if session.Values["FirstName"] != nil {
		cs.FirstName = session.Values["FirstName"].(string)
	}
	if session.Values["LastName"] != nil {
		cs.LastName = session.Values["LastName"].(string)
	}
	if session.Values["Username"] != nil {
		cs.Username = session.Values["Username"].(string)
	}
	if session.Values["DocID"] != nil {
		cs.DOCID = session.Values["DocID"].(string)
	}
	if session.Values["Avatar"] != nil {
		cs.Avatar = session.Values["Avatar"].(string)
	}
	err := templates.ExecuteTemplate(w, tmpl+".html", cs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
