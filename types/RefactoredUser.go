//New User Docs
package types

import (
	"time"

	"github.com/gorilla/sessions"
)

type UserLogin struct {
	Email                  string                  `json:"Email"`
	Password               []byte                  `json:"Password"`
	Token                  string                  `json:"Token"`
	Session                sessions.Session        `json:"Session"`
	LoginWarnings          map[string]LoginWarning `json:"LoginWarnings"`
	AuthenticationCode     string                  `json:"AuthenticationCode"`
	AuthenticationCodeTime time.Time               `json:"AuthenticationCodeTime"`
	FirstLogin             bool                    `json:"FirstLogin"`
}

type UserInfo struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	EmailVis  bool   `json:"EmailVis"`
	Gender    string `json:"Gender"`
	GenderVis bool   `json:"GenderVis"`
	Phone     string `json:"Phone"`
	PhoneVis  bool   `json:"PhoneVis"`

	Username string    `json:"Username"`
	Dob      time.Time `json:"Dob"`
}

type UserProfile struct {
	Avatar string `json:"Avatar"`

	Banner      string        `json:"Banner"`
	Projects    []ProjectInfo `json:"Projects"`
	Events      []EventInfo   `json:"Events"`
	Visible     bool          `json:"Visible"`
	Status      bool          `json:"Status"`
	QuickLinks  []Link        `json:"QuickLinks"`
	Tags        []string      `json:"Tags"`
	UserWidgets []string      `json:"UserWidgets"`
	Description []rune        `json:"Description"`
}

type FloatingHeadV2 struct {
	Avatar      string   `json:"Avatar"`
	FirstName   string   `json:"FirstName"`
	LastName    string   `json:"LastName"`
	Username    string   `json:"Username"`
	Tags        []string `json:"Tags"`
	Description []rune   `json:"Description"`
	University  string   `json:"UndergradSchool"`
}

type AccountInfo struct {
	Verified     bool      `json:"Verified"`
	AccType      int       `json:"AccType"` //highschool, college etc
	Organization string    `json:"Organization"`
	Category     string    `json:"Category"`
	Paid         bool      `json:"Paid"`
	AccCreation  time.Time `json:"AccCreation"`
}

type Education struct {
	HighSchool      string   `json:"HighSchool"`
	HSGradDate      string   `json:"HSGradDate"`
	CollegeGradDate string   `json:"CollegeGradDate"`
	University      string   `json:"UndergradSchool"`
	Majors          []string `json:"Majors"`
	Minors          []string `json:"Minors"`
	Class           int      `json:"Class"` //freshman:0,sophomore:1...
}

type Social struct {
	DocID            string          `json:"DocID"`
	UserFollowers    map[string]bool `json:"UserFollowers"`
	UserFollowing    map[string]bool `json:"UserFollowing"`
	ProjectFollowers map[string]bool `json:"ProjectFollowers"`
	ProjectFollowing map[string]bool `json:"ProjectFollowing"`
	EventFollowers   map[string]bool `json:"EventFollowers"`
	EventFollowing   map[string]bool `json:"EventFollowing"`
	Bell             map[string]bool `json:"Bell"`

	// Colleagues map[string]bool `json:"Colleagues"`
	// SentCollReq      []string `json:"SentCollReq"`
	// ReceivedCollReq  []string `json:"ReceivedCollReq"`

	BlockedUsers map[string]bool `json:"BlockedUsers"`
	BlockedBy    map[string]bool `json:"BlockedBy"`
}

type Joins struct {
	SentProjReq      map[string]bool `json:"SentProjReq"`
	ReceivedProjReq  map[string]bool `json:"ReceivedProjReq"`
	SentEventReq     map[string]bool `json:"SentEventReq"`
	ReceivedEventReq map[string]bool `json:"ReceivedEventReq"`
}

type Activity struct {
	EntryIDs []string `json:"EntryIDs"`
	// Entries               map[string]Entry                `json:"Entries"`
	SearchHist      []string `json:"SearchHist"`
	LikedEntryIDs   []string `json:"LikedEntryIds"`
	ProxyMessagesID string   `json:"ProxyMessagesID"`
}
