package client

import (
	htype "html/template"
	"net/http"

	sessions "github.com/gorilla/sessions"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//Eclient ... Reference to the ElasticSearch
var Eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

//Store ...
var Store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

var htmlPath = "/ustart/ustart_front/"

//ClientSide ... This struct represents a user state after he/she has logged in. Some fields may no longer be needed
//or are unnecessary.
type ClientSide struct {
	DOCID             string
	FirstName         string
	LastName          string
	Username          string
	Avatar            string
	ErrorOutput       error
	ErrorStatus       bool
	UserInfo          types.User
	Class             string
	Birthday          string
	ImageCode         string
	Description       string
	Followers         int
	UserFollowing     int
	ProjFollowing     int
	EventFollowing    int
	ProjectsFollowing int
	Page              string //DocID of current page
	FollowingStatus   bool
	ScrollID          string
	ListOfHeads       []types.FloatingHead
	ListOfHeads2      []types.FloatingHead
	ListOfHeads3      []types.FloatingHead
	Wall              []types.JournalEntry
	Widgets           []types.Widget
	Project           types.ProjectAggregate
	Event             types.EventAggregate
	Dashboard         types.DashboardAggregate
	Messages          []types.Message
	Hits              int
	Sent              string
}

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

	template := htype.Must(htype.ParseFiles(htmlPath + tmpl + ".html"))
	err := template.ExecuteTemplate(w, tmpl+".html", cs)
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

	template := htype.Must(htype.ParseFiles(htmlPath + tmpl + ".html"))
	err := template.ExecuteTemplate(w, tmpl+".html", cs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
