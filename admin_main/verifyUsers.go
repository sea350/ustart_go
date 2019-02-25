package main

import (
	getUser "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"

	// post "github.com/sea350/ustart_go/post/user"

	// admin "github.com/sea350/ustart_go/admin"

	"fmt"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	var usrs = []string{"fs817@nyu.edu", "ae1389@nyu.edu"}

	for _, usr := range usrs {
		// user, err := getUser.UserByEmail(eclient, usr)
		usrID, err := getUser.IDByUsername(eclient, usr)
		err = post.UpdateUser(eclient, usrID, "Verified", true)
		if err != nil {
			fmt.Println("LINE 24,", err)
		}
	}

}
