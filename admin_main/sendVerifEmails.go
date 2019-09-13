package main

import (
	// "github.com/sea350/ustart_go/uses"

	// admin "github.com/sea350/ustart_go/admin"

	"context"
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	ctx := context.Background()

	maq2 := elastic.NewBoolQuery()
	maq2 = maq.Must(elastic.NewTermQuery("Verified", false))
	maq2 = maq.Must(elastic.NewMatchQuery("Tags", "Coaching"))

	res2, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maq2).
		Size(500).
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
		emails = append(emails, data.Email)
	}

	notVer := int(res2.TotalHits())

	sendTo := []string{}
	fmt.Println("Unverified:", notVer)
	fmt.Println("The following users are unverified: ")
	for _, e := range emails {
		sendTo = append(sendTo, e)
		fmt.Println(e)
	}

	fmt.Println(sendTo)
	// for _, e := range sendTo {
	// 	uses.SendVerificationEmail(eclient, e)
	// }

}
