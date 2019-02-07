package main

import (
	"fmt"

	admin "github.com/sea350/ustart_go/admin"

	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {
	err := admin.ModifyBadge(eclient, "USTART", "give", "gl1144@nyu.edu", "")
	fmt.Println(err)
}
