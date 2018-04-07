package app_middleware

import (
	"errors"
	"fmt"
)

type response struct {
	Token   string `json:"Token"`
	Success bool   `json:"Success"`
	ErrMsg  error  `json:"ErrMsg"`
}

func setupResp() *response {
	resp := &response{
		Token:  "",
		ErrMsg: errors.New("Unknown error"),
	}

	return resp
}

func (r *response) updateResp(t string, em error, success bool) {

	r.Token = t
	r.ErrMsg = em
	r.Success = success

	if em != nil && em.Error() != "Password is incorrect" {
		fmt.Println("Error" + em.Error())
		panic(em)
	}
}
