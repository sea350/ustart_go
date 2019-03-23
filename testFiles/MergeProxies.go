package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	elastic "github.com/olivere/elastic"
	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
	postUser "github.com/sea350/ustart_go/post/user"
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
	ctx := context.Background()

	termQuery := elastic.NewTermQuery("DocID.keyword", userID)
	searchResult, err := client.Eclient.Search().
		Index(globals.ProxyMsgIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Proxies found: ", searchResult.TotalHits())
	var totalRemoved int
	var tempProxy types.ProxyMessages
	finalID := searchResult.Hits.Hits[len(searchResult.Hits.Hits)-1].Id
	masterList := make(map[string]types.ConversationState)
	if searchResult.TotalHits() > 1 {
		//do stuff
		for _, element := range searchResult.Hits.Hits {
			err := json.Unmarshal(*element.Source, &tempProxy)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				return
			}
			for _, convo := range tempProxy.Conversations {
				masterList[convo.ConvoID] = convo
			}
			if element.Id != finalID {
				if len(tempProxy.Conversations) == 0 {
					fmt.Println("empty proxy being removed")
				} else {
					fmt.Println("occupied proxy being removed ")
				}
				err = globals.DeleteByID(client.Eclient, element.Id, "proxymsg")
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					log.Println(err)
					return
				}
				totalRemoved++
				fmt.Println(totalRemoved)
			} else {
				fmt.Println(finalID)
				var lastArray []types.ConversationState
				for id := range masterList {
					_, err := getChat.ConvoByID(client.Eclient, id)
					if err == nil {
						lastArray = append(lastArray, masterList[id])
					}
				}
				err := postChat.UpdateProxyMsg(client.Eclient, finalID, "Conversations", lastArray)
				if err != nil {
					fmt.Println(lastArray)
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					log.Println(err)
					return
				}
			}
		}
		fmt.Println("Done!")
		err := postUser.UpdateUser(client.Eclient, userID, "ProxyMessagesID", finalID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}

	} else {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("no fixes needed")
	}

}
