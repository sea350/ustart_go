package types

import (
	"time"
)

//invite only, publicly viewable, whitelist, blacklist, settings, members, privileges, title, description, creation, event date, option for
//creator to cancel, visibility boolean, widgets, get/post, attending/invited/not going, members list and guest list
//Have it displayed on the profile
//People can search for the event based on title, url, tags, and PERSON HOSTING THE EVENT!
//Projects can host events and guests/members do the relation (projects can host and be the guest)

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
	GuestID        string        `json:"GuestID"`
	Representative []EventGuests `json:"Representative"` //For single user, nil by default
	Status         int           `json:"Status"`         //Marks whether they are invited/going/not going, 0 for invited, 1 for going, 2 for not
	Visible        bool          `json:"Visible"`
}

//EventMembers ... Member relevant data, who can modify the event page
type EventMembers struct {
	MemberID string    `json:"MemberID"`
	Role     int       `json:"Role"`
	JoinDate time.Time `json:"JoinDate"`
	Title    string    `json:"Title"`
	Visible  bool      `json:"Visible"`
}

//Events ... Event relevant data
type Events struct {
	Host              interface{}       `json:"Host"` //Displayed host on the event page, whether user or project
	Name              string            `json:"Name"`
	Tags              string            `json:"Tags"`
	Category          string            `json:"Category"`
	URLName           string            `json:"URLName"`
	QuickLinks        []Link            `json:"QuickLinks"`
	Description       []rune            `json:"Description"`
	Members           []EventMembers    `json:"Members"`
	Guests            []EventGuests     `json:"Guests"`
	EntryIDs          []string          `json:"EntryIDs"`
	Location          LocStruct         `json:"Location"`
	EventDateStart    time.Time         `json:"EventDateStart"`
	EventDateEnd      time.Time         `json:"EventDateEnd"`
	CreationDate      time.Time         `json:"CreationDate"`
	Widgets           []string          `json:"Widgets"`
	Whitelist         []string          `json:"Whitelist"`
	Blacklist         []string          `json:"Blacklist"`
	Status            bool              `json:"Status"` //Whether this event is still ongoing or cancelled
	Public            bool              `json:"Public"` //Whether this event is publicly viewable or invite-only viewable
	Avatar            string            `json:"Avatar"`
	CroppedAvatar     string            `json:"CropAvatar"`
	Banner            string            `json:"Banner"`
	Visible           bool              `json:"Visible"`
	MemberReqSent     []string          `json:"MemberReqSent"`
	MemberReqReceived []string          `json:"MemberReqReceived"`
	GuestReqSent      []string          `json:"GuestReqSent"`
	GuestReqReceived  []string          `json:"GuestReqReceived"`
	PrivilegeProfiles []EventPrivileges `json:"PrivilegeProfiles"`
}
