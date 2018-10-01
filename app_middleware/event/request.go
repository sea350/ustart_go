package event

import "github.com/sea350/ustart_go/backend/types"

type form struct {
	Username       string               `json:"Username"`
	EventID        string               `json:"EventID"`
	SessUser       types.AppSessionUser `json:"SessUser"`
	Representative []types.EventGuests  `json:"Representative"`
	Intent         string               `json:"Intent"`
	MemberJoinerID string               `json:"MemberJoinerID"`
	GuestJoinerID  string               `json:"GuestJoinerID"`
}

//Intent:
//join: toggles joining
