package signup

import (
	"errors"
	"fmt"
)

type response struct {
	Success bool  `json:"Success"`
	ErrMsg  error `json:"ErrMsg"`
}

func setupResp() *response {
	resp := &response{
		Success: false,
		ErrMsg:  errors.New("Unknown error"),
	}
	return resp
}

func (resp *response) updateResp(successful bool, err error) {
	resp.Success = successful
	resp.ErrMsg = err
	if err != nil {
		fmt.Println("Error in response for Signup")
	}
}
