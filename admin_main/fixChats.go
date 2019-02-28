package main

import (
	"context"
	"fmt"

	elastic "github.com/olivere/elastic"
	globals "github.com/sea350/ustart_go/globals"
)

//9v4r-GgBN3VvtvdieZzG
// g_5h42gBN3VvtvdiWZt3

//v4e02gBN3VvtvdiDZYs tarek doc id
var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {
	ctx := context.Background()

	query := elastic.NewTermQuery("_id", "-v4e02gBN3VvtvdiDZYs")

	res, err := eclient.Search().
		Index(globals.UserIndex).
		Query(query).
		Sort("_score", false).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Hits.TotalHits())
}
