package search

import (
	"fmt"

	types "github.com/sea350/ustart_go/types"
)

type response struct {
	Successful bool                 `json:"Successful"`
	Results    []types.FloatingHead `json:"Results"`
	ErrMsg     error                `json:"Error"`
}

func (r *response) update(success bool, em error, res []types.FloatingHead) {

	r.ErrMsg = em
	r.Successful = success
	r.Results = res

	if em != nil && em.Error() != "Password is incorrect" {
		fmt.Println("Error" + em.Error())
		panic(em)
	}
}
