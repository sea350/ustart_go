package types

import (
	"time"
)

//EventPrivileges ... Edit privileges for the Event Members
type EventPrivileges struct {
	RoleName     string `json:"RoleName"`
	RoleID       int    `json:"RoleID"`
	MemberManage bool   `json:"MemberEdit"`
	WidgetManage bool   `json:"WidgetManage"`
	PostManage   bool   `json:"PostManage"`
	Icon         bool   `json:"Icon"`
	Banner       bool   `json:"Banner"`
	Links        bool   `json:"Links"`
	Tags         bool   `json:"Tags"`
}

//EventGuests ... Guest relevant data, people who have been invited to attend the event
type EventGuests struct {
	GuestID        string `json:"GuestID"` //Can be a userID or a projectID
	Status         int    `json:"Status"`  //Marks whether they are invited/going/not going, 0 for invited, 1 for going, 2 for not
	Invisible      bool   `json:"Invisible"`
	Classification int    `json:"Classification"` //1 for guest, 2 for project guest
}

//EventMembers ... Member relevant data, who can modify the event page
type EventMembers struct {
	MemberID string    `json:"MemberID"`
	Role     int       `json:"Role"`
	JoinDate time.Time `json:"JoinDate"`
	Title    string    `json:"Title"`
	Visible  bool      `json:"Visible"`
}

//EventProjects ... Project information, what is displayed for the event page
type EventProjects struct {
	ProjectID      string        `json:"ProjectID"`
	Title          string        `json:"Title"`
	Visible        bool          `json:"Visible"`
	Representative []EventGuests `json:"Representative"`
}

//Events ... Event relevant data
type Events struct {
	EventID           string            `json:"EventID"`
	Host              string            `json:"Host"` //Displayed host on the event page, by userID or projectID
	IsProjectHost     bool              `json:"IsProjectHost"`
	Name              string            `json:"Name"`
	Tags              []string          `json:"Tags"`
	Category          string            `json:"Category"`
	URLName           string            `json:"URLName"`
	FollowedUsers     []string          `json:"FollowedUsers"`
	QuickLinks        []Link            `json:"QuickLinks"`
	Description       []rune            `json:"Description"`
	Members           []EventMembers    `json:"Members"`
	Projects          []EventProjects   `json:"Projects"`
	Guests            []EventGuests     `json:"Guests"`
	EntryIDs          []string          `json:"EntryIDs"`
	Location          LocStruct         `json:"Location"`
	Organization      string            `json:"Organization"`
	EventDateStart    time.Time         `json:"EventDateStart"`
	EventDateEnd      time.Time         `json:"EventDateEnd"`
	CreationDate      time.Time         `json:"CreationDate"`
	Widgets           []string          `json:"Widgets"`
	Status            bool              `json:"Status"` //Whether this event is still ongoing or cancelled
	Avatar            string            `json:"Avatar"`
	CroppedAvatar     string            `json:"CropAvatar"`
	Banner            string            `json:"Banner"`
	Visible           bool              `json:"Visible"`
	MemberReqSent     []string          `json:"MemberReqSent"`
	MemberReqReceived []string          `json:"MemberReqReceived"`
	GuestReqSent      map[string]int    `json:"GuestReqSent"`
	GuestReqReceived  map[string]int    `json:"GuestReqReceived"`
	PrivilegeProfiles []EventPrivileges `json:"PrivilegeProfiles"`
}
