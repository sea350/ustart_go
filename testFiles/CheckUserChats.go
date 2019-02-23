package main

import (
	"context"
	"fmt"
	"strings"

	elastic "github.com/olivere/elastic"
	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"
)

func main() {
	fmt.Println("getting user by username")
	id, err := get.IDByUsername(client.Eclient, "th1750")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User docID: " + id)
	usr, err := get.UserByID(client.Eclient, id)
	fmt.Println(usr.ProxyMessagesID)
	fmt.Println(id == "-v4e02gBN3VvtvdiDZYs")
	proxyid, err := getChat.ProxyIDByUserID(client.Eclient, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User msg proxies: " + proxyid)
	id = "v4e02gBN3VvtvdiDZYs"
	/*
		proxy, err := getChat.ProxyMsgByID(client.Eclient, proxyid)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Printing cached convos: ")
		for i := range proxy.Conversations {
			fmt.Println(i)
			fmt.Println(proxy.Conversations[i].ConvoID)
		}
	*/
	query := elastic.NewBoolQuery()

	query = query.Must(elastic.NewTermQuery("Eavesdroppers.DocID", strings.ToLower(id)))

	fmt.Println("Printing queried convos: ")
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
		fmt.Println("--------------------------------")
		chat, err := getChat.ConvoByID(client.Eclient, hit.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(hit.Id)
		fmt.Println(chat.ReferenceID)
		fmt.Println(chat.Eavesdroppers)
		fmt.Println(chat.Class)
		fmt.Println(chat.Size)
	}
}
