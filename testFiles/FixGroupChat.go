package main

import (
	"context"
	"fmt"
	"strings"

	getChat "github.com/sea350/ustart_go/get/chat"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"
	elastic "github.com/olivere/elastic"
)

func main() {
	// fmt.Println("getting user by username")
	// id, err := get.IDByUsername(client.Eclient, "th1750") //
	// id, err := get.IDByUsername(client.Eclient, "support")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("User docID: " + id)
	// proxyid, err := getChat.ProxyIDByUserID(client.Eclient, id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(proxyid)
	// proxy, err := getChat.ProxyMsgByID(client.Eclient, proxyid)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(proxy)

	// var chatFound bool
	// for i := range proxy.Conversations {
	// 	if proxy.Conversations[i].ConvoID == "9v4r-GgBN3VvtvdieZzG" {
	// 		proxy.Conversations = append(proxy.Conversations[:i], proxy.Conversations[i+1:]...)
	// 		fmt.Println(len(proxy.Conversations))
	// 		fmt.Println(i)
	// 		chatFound = true
	// 		break
	// 	}
	// }
	// fmt.Println("chat found status: ", chatFound)

	// if chatFound {
	// 	err = postChat.UpdateProxyMsg(client.Eclient, proxyid, "Conversations", proxy.Conversations)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	query := elastic.NewBoolQuery()

	query = query.Must(elastic.NewTermQuery("ReferenceID", strings.ToLower("9_6_6WgBN3VvtvdiTJsk")))

	ctx := context.Background() //intialize context background
	searchResults, err := client.Eclient.Search().
		Index(globals.ConvoIndex).
		Query(query).
		Pretty(true).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}
	if searchResults.TotalHits() == 0 {
		fmt.Println("empty")
		return
	} else if searchResults.TotalHits() > 1 {
		fmt.Println("That's not supposed to happen")
		return
	}
	for _, hit := range searchResults.Hits.Hits {
		chat, err := getChat.ConvoByID(client.Eclient, hit.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(chat.Eavesdroppers)
	}
}
