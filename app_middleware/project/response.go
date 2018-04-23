package project

import (
	"fmt"

	"github.com/sea350/ustart_go/types"
)

type response struct {
	Successful bool          `json:"Successful"`
	Project    types.Project `json:"Project"`
	ErrMsg     error         `json:"Error"`
}

func (r *response) update(success bool, em error, proj types.Project) {

	r.ErrMsg = em
	r.Successful = success
	r.Project = proj

	if em != nil && em.Error() != "Password is incorrect" {
		fmt.Println("Error" + em.Error())
		panic(em)
	}
}
