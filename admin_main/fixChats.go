package main

import (
	"context"
	"fmt"
	"strings"

	elastic "github.com/olivere/elastic"
	globals "github.com/sea350/ustart_go/globals"
)

//9v4r-GgBN3VvtvdieZzG
// g_5h42gBN3VvtvdiWZt3

//v4e02gBN3VvtvdiDZYs tarek doc id
var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {
	ctx := context.Background()

	query := elastic.NewTermQuery("PosterID", strings.ToLower("v4e02gBN3VvtvdiDZYs"))

	res, err := eclient.Search().
		Index(globals.EntryIndex).
		Query(query).
		Sort("_score", false).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.TotalHits())
}
