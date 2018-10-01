package event

import (
	"fmt"

	"github.com/sea350/ustart_go/backend/types"
)

type response struct {
	Successful bool         `json:"Successful"`
	Event      types.Events `json:"Event"`
	ErrMsg     error        `json:"Error"`
}

func (r *response) update(success bool, em error, evnt types.Events) {

	r.ErrMsg = em
	r.Successful = success
	r.Event = evnt

	if em != nil && em.Error() != "Password is incorrect" {
		fmt.Println("Error" + em.Error())
		panic(em)
	}
}
