package login

import (
	"errors"
	"fmt"
)

type response struct {
	//Token   string            `json:"Token"`
	Success bool  `json:"Success"`
	ErrMsg  error `json:"ErrMsg"`
	//SessUsr *sessions.Session `json:"SessUsr"`
}

func setupResp() *response {
	resp := &response{
		Success: true,
		ErrMsg:  errors.New("Unknown error"),
	}

	return resp
}

func (r *response) updateResp(em error, success bool) {

	r.ErrMsg = em
	r.Success = success

	if em != nil && em.Error() != "Password is incorrect" {
		fmt.Println("Error" + em.Error())
		panic(em)
	}
}
