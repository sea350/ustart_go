package main

import (
	// getUser "github.com/sea350/ustart_go/get/user"
	// post "github.com/sea350/ustart_go/post/user"

	get "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"

	// post "github.com/sea350/ustart_go/post/user"

	// admin "github.com/sea350/ustart_go/admin"

	"fmt"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	proj, err := get.ProjectByURL(eclient, "pottymints")
	if err != nil {
		fmt.Println(err)
	}

	reqs := proj.MemberReqReceived

	if len(reqs) > 0 {
		fmt.Println("Number of requests:", len(reqs))
		for _, id := range reqs {
			usr, err := getUser.UserByID(eclient, id)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(usr.FirstName, usr.LastName, usr.Username)

		}
	} else {
		fmt.Println("No member join requests")
	}

}
