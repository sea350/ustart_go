package types

import (
	"time"
)

type Member struct {
	MemberID string    `json:MemberID`
	Role     int       `json:Role`
	JoinDate time.Time `json:JoinDate`
	Title    string    `json:Title`
	Visible  bool      `json:Visible`
}

type Project struct {
	Name              string    `json:Name`
	URLName           string    `json:URLName`
	Members           []Member  `json:Members`
	Location          LocStruct `json:Location`
	EntryIDs          []string  `json:EntryIDs`
	Category          string    `json:Category`
	ListNeeded        []string  `json:ListNeeded`
	CreationDate      time.Time `json:CreationDate`
	Visible           bool      `json:Visible`
	Status            bool      `json:Status`
	QuickLinks        []Link    `json:QuickLinks`
	Avatar            string    `json:Avatar`
	CroppedAvatar     string    `json:CropAvatar`
	Banner            string    `json:Banner`
	Description       []rune    `json:Description`
	Tags              []string  `json:Tags`
	BlockedUsers      []string  `json:BlockedUsers`
	ConversationIDs   []string  `json:ConversationIDs`
	MemberReqSent     []string  `json:MemberReqSent`
	MemberReqReceived []string  `json:MemberReqReceived`
	Widgets           []string  `json:Widgets`
}
