package main

import (

	// admin "github.com/sea350/ustart_go/admin"

	"context"
	"encoding/json"
	"fmt"

	"github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
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
		Query(maq2).
		Size(100).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	emails := []string{}
	for _, id := range res2.Hits.Hits {
		data := types.User{}
		err = json.Unmarshal(*id.Source, &data)
		if err != nil {
			fmt.Println(err)
		}
		emails = append(email, data.Email)
	}
	ver := int(res.TotalHits())
	notVer := int(res2.TotalHits())
	fmt.Println("Total:", ver+notVer)
	fmt.Println("Verified:", ver)
	fmt.Println("Unverified:", notVer)
	fmt.Println("The following users are unverified: ")
	for _, e := range emails {
		fmt.Println(e)
	}

}
