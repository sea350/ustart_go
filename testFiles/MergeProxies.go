package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	elastic "github.com/olivere/elastic"
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

	proxiesRemaining := searchResult.TotalHits()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Proxies found: ", proxiesRemaining)
	var totalRemoved int
	var tempProxy types.ProxyMessages
	var finalID string
	masterList := make(map[string]types.ConversationState)
	if proxiesRemaining > 1 {
		//do stuff
		for _, element := range searchResult.Hits.Hits {
			err := json.Unmarshal(*element.Source, &tempProxy)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				return
			}
			if proxiesRemaining > 1 {
				if len(tempProxy.Conversations) == 0 {
					err = globals.DeleteByID(client.Eclient, element.Id, "proxymsg")
					if err != nil {
						log.SetFlags(log.LstdFlags | log.Lshortfile)
						log.Println(err)
						return
					}
					totalRemoved++
					fmt.Println(totalRemoved)
				} else {
					for _, convo := range tempProxy.Conversations {
						masterList[convo.ConvoID] = convo
					}
					err = globals.DeleteByID(client.Eclient, element.Id, "proxymsg")
					if err != nil {
						log.SetFlags(log.LstdFlags | log.Lshortfile)
						log.Println(err)
						return
					}
					totalRemoved++
					fmt.Println(totalRemoved)
				}
			} else {
				finalID = element.Id
				fmt.Println(finalID)
				for _, convo := range tempProxy.Conversations {
					masterList[convo.ConvoID] = convo
				}
				var lastArray []types.ConversationState
				for id := range masterList {
					lastArray = append(lastArray, masterList[id])
				}
				err := postChat.UpdateProxyMsg(client.Eclient, finalID, "Conversations", lastArray)
				if err != nil {
					fmt.Println(lastArray)
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					log.Println(err)
					return
				}
			}
			proxiesRemaining--

		}
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
