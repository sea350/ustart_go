package profile

import "github.com/sea350/ustart_go/types"

type form struct {
	Username string               `json:"Username"`
	SessUser types.AppSessionUser `json:"SessUser"`
	Intent   string               `json:"Intent"`
}

//Intent:
//uf: user follow
//uu : user unfollow
