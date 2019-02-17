package main

import (
	"context"
	"fmt"
	"strings"

	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"
	elastic "gopkg.in/olivere/elastic.v5"
)

func main() {
	fmt.Println("getting user by email")
	// id, err := get.IDByUsername(client.Eclient, "th1750")
	id, err := get.IDByUsername(client.Eclient, "HeatherMT")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User docID: " + id)
	proxyid, err := getChat.ProxyIDByUserID(client.Eclient, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(proxyid)
	proxy, err := getChat.ProxyMsgByID(client.Eclient, proxyid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(proxy)

	query := elastic.NewBoolQuery()

	query = query.Must(elastic.NewTermQuery("Eavesdroppers.DocID", strings.ToLower(id)))

	ctx := context.Background() //intialize context background
	searchResults, err := client.Eclient.Search().
		Index(globals.ConvoIndex).
		Query(query).
		Pretty(true).
		Do(ctx)

	if searchResults.TotalHits() == 0 {
		fmt.Println("empty")
		return
	}
	for _, hit := range searchResults.Hits.Hits {
		chat, err := getChat.ConvoByID(client.Eclient, hit.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(chat)
	}
}
