package login

import (
	"errors"
	"fmt"

	"github.com/sea350/ustart_go/types"
)

type response struct {
	//Token   string            `json:"Token"`
	Success     bool                 `json:"Success"`
	ErrMsg      error                `json:"ErrMsg"`
	SessionUser types.AppSessionUser `json:"SessionUser"`
}

func setupResp() *response {
	resp := &response{
		Success: true,
		ErrMsg:  errors.New("Unknown error"),
	}

	return resp
}

func (r *response) updateResp(em error, success bool, sessUsr types.AppSessionUser) {

	r.ErrMsg = em
	r.Success = success
	r.SessionUser = sessUsr

	if em != nil && em.Error() != "Password is incorrect" {
		fmt.Println("Error" + em.Error())
		panic(em)
	}
}
