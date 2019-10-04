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

	// _, err := getUser.UserByUsername(eclient, "support")
	usrID, err := getUser.IDByUsername(eclient, "support")

	// fmt.Println(len(usr.Projects))
	// updatesProj := append(usr.Projects[:3], usr.Projects[4:]...)
	err = post.UpdateUser(eclient, usrID, "Email", "support@ustart.today")
	if err != nil {
		fmt.Println("LINE 24,", err)
	}

}
