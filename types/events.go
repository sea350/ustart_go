package types

import (
	"time"
)

//invite only, publicly viewable, whitelist, blacklist, settings, members, privileges, title, description, creation, event date, option for
//creator to cancel, visibility boolean, widgets, get/post, attending/invited/not going, members list and guest list
//Have it displayed on the profile
//People can search for the event based on title, url, tags, and PERSON HOSTING THE EVENT!

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
	GuestID string `json:"GuestID"`
	Status  int    `json:"Status"` //Marks whether they are invited/going/not going, 0 for invited, 1 for going, 2 for not
	Visible bool   `json:"Visible"`
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
	Host              User              `json:"Host"`
	Name              string            `json:"Name"`
	Category          string            `json:"Category"`
	URLName           string            `json:"URLName"`
	Description       []rune            `json:"Description"`
	Members           []EventMembers    `json:"Members"`
	Guests            []EventGuests     `json:"Guests"`
	Location          LocStruct         `json:"Location"`
	EventDate         time.Time         `json:"EventDate"`
	CreationDate      time.Time         `json:"CreationDate"`
	Widgets           []Widget          `json:"Widgets"`
	Whitelist         []string          `json:"Whitelist"`
	Blacklist         []string          `json:"Blacklist"`
	Status            bool              `json:"Status"` //Whether this event is still ongoing or cancelled
	Public            bool              `json:"Public"` //Whether this event is publicly viewable or invite-only viewable
	Avatar            string            `json:"Avatar"`
	CroppedAvatar     string            `json:"CropAvatar"`
	Banner            string            `json:"Banner"`
	MemberReqSent     []string          `json:"MemberReqSent"`
	PrivilegeProfiles []EventPrivileges `json:"PrivilegeProfiles"`
	//MemberReqReceived []string       `json:"MemberReqReceived"` We probably don't want this
}