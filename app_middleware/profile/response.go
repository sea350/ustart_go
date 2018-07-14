package profile

import (
	"fmt"

	"github.com/sea350/ustart_go/types"
)

type response struct {
	Successful bool       `json:"Successful"`
	User       types.User `json:"User"`
	ErrMsg     error      `json:"Error"`
	CreatedID  string     `json:"CreatedID"`
}

func (r *response) update(success bool, em error, created string, usr types.User) {

	r.ErrMsg = em
	r.Successful = success
	r.User = usr
	r.CreatedID = created

	if em != nil && em.Error() != "Password is incorrect" {
		fmt.Println("Error" + em.Error())
		panic(em)
	}
}
