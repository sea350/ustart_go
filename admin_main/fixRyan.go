package main

import (
	getUser "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"

	// post "github.com/sea350/ustart_go/post/user"

	// admin "github.com/sea350/ustart_go/admin"

	"fmt"

	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	usr, err := getUser.UserByUsername(eclient, "ryanrozbiani")
	usrID, err := getUser.IDByUsername(eclient, "ryanrozbiani")

	fmt.Println(len(usr.Projects))
	updatesProj := append(usr.Projects[:3], usr.Projects[4:]...)
	err = post.UpdateUser(eclient, usrID, "Projects", updatesProj)
	if err != nil {
		fmt.Println("LINE 24,", err)
	}

	fmt.Println(len(updatesProj))

}
