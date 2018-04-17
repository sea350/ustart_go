package settings

import (
	"github.com/gorilla/sessions"
)

type response struct {
	Successful bool              `json:"Successful"`
	ErrMsg     error             `json:"ErrMsg"`
	Retreived  string            `json:"Retreived"`
	SessUsr    *sessions.Session `json:"SessUsr"`
	//ColorPalette string `json:"ColorPalette"`
}

func (r *response) update(s bool, e error, ret string, sess *sessions.Session) {
	r.Successful = s
	r.ErrMsg = e
	r.SessUsr = sess
	if e != nil {
		panic(e)
	}
}
