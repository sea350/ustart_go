package main

import (

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

	maq := elastic.NewBoolQuery()
	maq = maq.Must(elastic.NewTermQuery("Verified", true))
	maq = maq.Must(elastic.NewMatchQuery("Tags", "Coaching"))

	res, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maq).
		Size(500).
		// Sort("_score", true).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Printing emails: ")
	for _, id := range res.Hits.Hits {
		data := types.User{}
		err = json.Unmarshal(*id.Source, &data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data.Email)
	}

}
