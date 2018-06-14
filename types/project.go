package types

import (
	"time"
)

type Privileges struct {
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

//Member ... all member relevant information
type Member struct {
	MemberID string    `json:"MemberID"`
	Role     int       `json:"Role"`
	JoinDate time.Time `json:"JoinDate"`
	Title    string    `json:"Title"`

	Visible bool `json:"Visible"`
}

//Project ... Project relevant data
type Project struct {
	Name              string       `json:"Name"`
	URLName           string       `json:"URLName"`
	Members           []Member     `json:"Members"`
	Location          LocStruct    `json:"Location"`
	EntryIDs          []string     `json:"EntryIDs"`
	Category          string       `json:"Category"`
	ListNeeded        []string     `json:"ListNeeded"`
	CreationDate      time.Time    `json:"CreationDate"`
	Visible           bool         `json:"Visible"`
	Status            bool         `json:"Status"`
	QuickLinks        []Link       `json:"QuickLinks"`
	Avatar            string       `json:"Avatar"`
	CroppedAvatar     string       `json:"CropAvatar"`
	Banner            string       `json:"Banner"`
	Description       []rune       `json:"Description"`
	Tags              []string     `json:"Tags"`
	BlockedUsers      []string     `json:"BlockedUsers"`
	ConversationIDs   []string     `json:"ConversationIDs"`
	MemberReqSent     []string     `json:"MemberReqSent"`
	MemberReqReceived []string     `json:"MemberReqReceived"`
	Organization      string       `json:"Organization"`
	Widgets           []string     `json:"Widgets"`
	PrivilegeProfiles []Privileges `json:"PrivilegeProfiles"`
	FollowedUsers     []string     `json:"FollowedUsers"`
}
