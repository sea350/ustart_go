package types

import (
	"time"
)

//WARNING: NOT FOR DATABASE USE

//AppSessionUser ...
type AppSessionUser struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	DocID     string `json:"DocID"`
}

//SessionUser ... All data needed to be stored in session
type SessionUser struct {
	FirstName string      `json:"FirstName"`
	LastName  string      `json:"LastName"`
	Username  string      `json:"Username"`
	Email     string      `json:"Email"`
	DocID     string      `json:"ID"`
	Avatar    string      `json:"Avatar"`
	Interface interface{} `json:"Interface"`
}

//FloatingHead ... All data needed for a generic user icon
type FloatingHead struct {
	Username       string      `json:"Username"`
	FirstName      string      `json:"FirstName"`
	LastName       string      `json:"LastName"`
	Email          string      `json:"Email"`
	Tags           []string    `json:"Tags"`
	Image          string      `json:"Image"`
	Followed       bool        `json:"Followed"`
	Bio            []rune      `json:"Bio"`
	DocID          string      `json:"DocID"`
	Classification int         `json:"Classification"`
	Notifications  int         `json:"Notifications"`
	Time           time.Time   `json:"Time"`
	Read           bool        `json:"Read"`
	Interface      interface{} `json:"Interface"`
	//for projects Username = project URLName
	//for projects Firstname = project Name
	//for projects Bio = project Description
}

//JournalEntry ... All data needed to display an entry
type JournalEntry struct {
	ElementID        string      `json:"ElementID"`
	FirstName        string      `json:"FirstName"`
	LastName         string      `json:"LastName"`
	Image            string      `json:"Image"`
	Element          Entry       `json:"Element"`
	NumReplies       int         `json:"NumReplies"`
	NumLikes         int         `json:"NumLikes"`
	NumShares        int         `json:"NumShares"`
	Liked            bool        `json:"Liked"`
	ReferenceElement interface{} `json:"ReferenceElement"`
}

//ProjectSubWidgets ... data specific to each project being displayed on the projects widget
type ProjectSubWidgets struct {
	Name   string `json:"Name"`
	Link   string `json:"Link"`
	Avatar string `json:"Avatar"`
	ID     string `json:"ID"`
}

//ProjectAggregate ... a compiled struct of all relevant project data
type ProjectAggregate struct {
	DocID          string         `json:"DocID"`
	ProjectData    Project        `json:"ProjectData"`
	MemberData     []FloatingHead `json:"MemberData"`
	Editable       bool           `json:"Editable"`
	RequestAllowed bool           `json:"RequestAllowed"`
}

//DashboardAggregate ... a compiled struct of all relevant dashboard data
type DashboardAggregate struct {
	DocID          string    `json:"DocID"`
	DashboardData  Dashboard `json:"DashboardData"`
	RequestAllowed bool      `json:"RequestAllowed"`
}

//EventSubWidgets ... data specific to each event being displayed on the events widget
type EventSubWidgets struct {
	Name   string `json:"Name"`
	Link   string `json:"Link"`
	Avatar string `json:"Avatar"`
	ID     string `json:"ID"`
}

//EventAggregate ... a compiled struct of all relevant event data
type EventAggregate struct {
	DocID          string         `json:"DocID"`
	EventData      Events         `json:"EventData"`
	MemberData     []FloatingHead `json:"MemberData"`
	GuestData      []FloatingHead `json:"GuestData"`
	ProjectData    []FloatingHead `json:"ProjectData"`
	Editable       bool           `json:"Editable"`
	RequestAllowed bool           `json:"RequestAllowed"`
}
