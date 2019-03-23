package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	elastic "github.com/olivere/elastic"
	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
	types "github.com/sea350/ustart_go/types"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter user's username: ")
	username1, _ := reader.ReadString('\n')
	username1 = username1[:len(username1)-1]
	userID, err := get.IDByUsername(client.Eclient, username1)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}
	fmt.Println("User docID: " + userID)

	proxyid, err := getChat.ProxyIDByUserID(client.Eclient, userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User msg proxies: " + proxyid)

	ctx := context.Background()

	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("Eavesdroppers.DocID.keyword", userID))
	// query = query.Must(elastic.NewTermQuery("Class", 1))

	searchResult, err := client.Eclient.Search().
		Index(globals.ConvoIndex).
		Query(query).
		Do(ctx)

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Chats found: ", searchResult.TotalHits())

	var lastArray []types.ConversationState
	for _, element := range searchResult.Hits.Hits {
		err := postChat.AppendToProxy(client.Eclient, proxyid, element.Id, true)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}
	}

}
