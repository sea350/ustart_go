package main

import (
	getFollow "github.com/sea350/ustart_go/get/follow"
	getUser "github.com/sea350/ustart_go/get/user"

	// post "github.com/sea350/ustart_go/post/user"
	postFollow "github.com/sea350/ustart_go/post/follow"

	// admin "github.com/sea350/ustart_go/admin"

	"fmt"

	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	ryanID, err := getUser.IDByUsername(eclient, "ryanrozbiani")
	if err != nil {
		fmt.Println("LINE 24,", err)
	}

	follID, foll, err := getFollow.ByID(eclient, ryanID)
	if err != nil {
		fmt.Println("LINE29,", err)
	}

	var emptyMap = make(map[string]bool)
	foll.ProjectFollowing = emptyMap

	err = postFollow.ReindexFollow(eclient, follID, foll)
	if err != nil {
		fmt.Println("LINE 40,", err)
	}
}
