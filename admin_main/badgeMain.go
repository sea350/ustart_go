package main

import (
	"fmt"

	// admin "github.com/sea350/ustart_go/admin"

	get "github.com/sea350/ustart_go/get/badge"
	"github.com/sea350/ustart_go/globals"
	post "github.com/sea350/ustart_go/post/badge"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {
	// err := admin.ModifyBadge(eclient, "USTART", "give", "gl1144@nyu.edu", "")
	badge, err := get.BadgeByType(eclient, "USTART")

	err = post.UpdateBadge(eclient, badge.ID, "Roster", append(badge.Roster, "gl1144@nyu.edu"))
	fmt.Println(err)
}
