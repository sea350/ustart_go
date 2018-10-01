package project

import "github.com/sea350/ustart_go/backend/types"

type form struct {
	Username  string               `json:"Username"`
	ProjectID string               `json:"ProjectID"`
	SessUser  types.AppSessionUser `json:"SessUser"`
	Intent    string               `json:"Intent"`
	JoinerID  string               `json:"JoinerID"`
}

//Intent:
//join: toggles joining
