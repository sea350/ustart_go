package main

import (
	"fmt"

	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
)

func main() {
	fmt.Println("getting user by username")
	// id, err := get.IDByUsername(client.Eclient, "th1750") //
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

	var chatFound bool
	for i := range proxy.Conversations {
		if proxy.Conversations[i].ConvoID == "9v4r-GgBN3VvtvdieZzG" {
			if i == len(proxy.Conversations)-1 {
				proxy.Conversations = proxy.Conversations[:i]
			} else {
				proxy.Conversations = append(proxy.Conversations[:i], proxy.Conversations[i+1:]...)
			}
			chatFound = true
		}
	}
	fmt.Println("chat found status: ", chatFound)

	if chatFound {
		err = postChat.UpdateProxyMsg(client.Eclient, proxyid, "Conversations", proxy.Conversations)
		if err != nil {
			fmt.Println(err)
		}
	}

	// query := elastic.NewBoolQuery()

	// query = query.Must(elastic.NewTermQuery("Eavesdroppers.DocID", strings.ToLower(id)))

	// ctx := context.Background() //intialize context background
	// searchResults, err := client.Eclient.Search().
	// 	Index(globals.ConvoIndex).
	// 	Query(query).
	// 	Pretty(true).
	// 	Do(ctx)

	// if searchResults.TotalHits() == 0 {
	// 	fmt.Println("empty")
	// 	return
	// }
	// for _, hit := range searchResults.Hits.Hits {
	// 	chat, err := getChat.ConvoByID(client.Eclient, hit.Id)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	fmt.Println(hit.Id, chat.ReferenceID, chat.Class, chat.Size)
	// }
}
