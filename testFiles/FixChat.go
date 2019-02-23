package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"

	"github.com/sea350/ustart_go/middleware/client"
)

func main() {
	ctx := context.Background()

	termQuery := elastic.NewTermQuery("DocID", strings.ToLower("v4e02gBN3VvtvdiDZYs"))
	searchResult, err := client.Eclient.Search().
		Index(globals.ProxyMsgIndex).
		Query(termQuery).
		Do(ctx)

	var proxyID string

	for i, element := range searchResult.Hits.Hits {
		if "Bf4g02gBN3VvtvdiF5fV" != element.Id {
			err := globals.DeleteByID(client.Eclient, element.Id, "proxymsg")
			if err != nil {
				fmt.Println(element.Id + "failed to be deleted")
				fmt.Println(err)
			} else {
				fmt.Println("number of proxies deleted = " + strconv.Itoa(i))
			}
		}

	}

}
