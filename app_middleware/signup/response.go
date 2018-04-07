package signup

import "errors"

type response struct {
	Successful bool  `json:"Successful"`
	ErrMsg     error `json:"ErrMsg"`
}

func setupResp() *response {
	resp := &response{
		Successful: false,
		ErrMsg:     errors.New("Unknown error"),
	}
	return resp
}

func (resp *response) updateResp(successful bool, err error) {
	resp.Successful = successful
	resp.ErrMsg = err
	if err != nil {
		panic(err)
	}
}
