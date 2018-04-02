package app_middleware

import (
	"errors"
	"fmt"
)

type response struct {
	Token  string `json:Token`
	ErrMsg error  `json:ErrMsg`
}

func setupResp() *response {
	resp := &response{
		Token:  "",
		ErrMsg: errors.New("Unknown error"),
	}

	return resp
}

func (r *response) updateResp(t string, em error) {

	r.Token = t
	r.ErrMsg = em

	if em != nil && em.Error() != "Password is incorrect" {
		fmt.Println("Error" + em.Error())
		panic(em)
	}
}
