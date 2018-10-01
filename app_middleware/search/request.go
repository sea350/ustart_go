package search

import "github.com/sea350/ustart_go/backend/types"

type form struct {
	Term     string               `json:"Term"`
	SessUser types.AppSessionUser `json:"SessUser"`
	Intent   string               `json:"Intent"`
}

//Intent:
//gen: search everything
//usr : search users
//proj : search projects
//skill : search skills
