package main

import (

	// admin "github.com/sea350/ustart_go/admin"

	"context"
	"fmt"

	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	ctx := context.Background()

	maq := elastic.NewTermQuery("Verified", true)
	res, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maq).
		Size(100).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	maq2 := elastic.NewTermQuery("Verified", false)
	res2, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maq).
		Size(100).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Total:", res.Hits.TotalHits()+res2.Hits.TotalHits())
	fmt.Println("Verified:", res.Hits.TotalHits())
	fmt.Println("Unverified:", res2.Hits.TotalHits())

}
