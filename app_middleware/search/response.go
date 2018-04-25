package search

import (
	"fmt"
)

type response struct {
	Successful bool                 `json:"Successful"`
	Results    []FloatingSearchHead `json:"Results"`
	ErrMsg     error                `json:"Error"`
}

func (r *response) update(success bool, em error, res []FloatingSearchHead) {

	r.ErrMsg = em
	r.Successful = success
	r.Results = res

	if em != nil && em.Error() != "Password is incorrect" {
		fmt.Println("Error" + em.Error())
		panic(em)
	}
}
