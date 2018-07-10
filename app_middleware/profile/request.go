package profile

import (
	"time"

	"github.com/sea350/ustart_go/types"
)

type form struct {
	Username    string               `json:"Username"`
	SessUser    types.AppSessionUser `json:"SessUser"`
	Title       string               `json:"Title"`
	Description string               `json:"Description"`
	CustomURL   string               `json:"CustomURL"`
	Category    string               `json:"Category"`
	Location    types.LocStruct      `json:"Location"`
	EventStart  time.Time            `json:"EventStart"`
	EventEnd    time.Time            `json:"EventEnd"`

	Intent string `json:"Intent"`
}

//Intent:
//uf: user follow
//uu : user unfollow
