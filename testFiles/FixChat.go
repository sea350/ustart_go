package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	elastic "github.com/olivere/elastic"
	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/chat"
	types "github.com/sea350/ustart_go/types"
)

func main() {
	fmt.Println("getting user by username")
	id, err := get.IDByUsername(client.Eclient, "AkbarMalikov")
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
	var dash = rune('-')
	var underscore = rune('_')
	var tempRuneArr []rune
	for _, char := range id {
		if char != dash && char != underscore {
			tempRuneArr = append(tempRuneArr, char)
		}
	}
	trimmedUser1 := string(tempRuneArr)

	fmt.Println("User msg proxies: " + proxyid)

	proxy, err := getChat.ProxyMsgByID(client.Eclient, proxyid)
	if err != nil {
		fmt.Println(err)
		return
	}
	//-------------------------------------------
	fmt.Println("getting user by username")
	id2, err := get.IDByUsername(client.Eclient, "min")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User docID: " + id2)
	proxyid2, err := getChat.ProxyIDByUserID(client.Eclient, id2)
	if err != nil {
		fmt.Println(err)
		return
	}
	// tempRuneArr = []rune{}
	// for _, char := range id2 {
	// 	if char != dash && char != underscore {
	// 		tempRuneArr = append(tempRuneArr, char)
	// 	}
	// }
	// trimmedUser2 := string(tempRuneArr)

	fmt.Println("User msg proxies: " + proxyid2)

	proxy2, err := getChat.ProxyMsgByID(client.Eclient, proxyid)
	if err != nil {
		fmt.Println(err)
		return
	}
	//-------------------------------------

	fmt.Println("Printing cached convos: ")
	for i := range proxy.Conversations {
		fmt.Println(i)
		fmt.Println(proxy.Conversations[i].ConvoID)
	}

	query := elastic.NewBoolQuery()

	query = query.Must(elastic.NewTermQuery("Eavesdroppers.DocID", strings.ToLower(trimmedUser1)))
	query = query.Must(elastic.NewTermQuery("Eavesdroppers.DocID", strings.ToLower(id2)))
	query = query.Must(elastic.NewTermQuery("Class", "1"))

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

	for i, hit := range searchResults.Hits.Hits {
		fmt.Println("--------------------------------")
		// chat, err := getChat.ConvoByID(client.Eclient, hit.Id)
		// if err != nil {
		// 	fmt.Println(err)
		// 	continue
		// }
		fmt.Println(hit.Id)
		// fmt.Println(chat.ReferenceID)
		// fmt.Println(chat.Eavesdroppers)
		// fmt.Println(chat.Class)
		// fmt.Println(chat.Size)

		tempProxyConvos1 := []types.ConversationState{}
		for _, elem := range proxy.Conversations {
			if elem.ConvoID != hit.Id {
				tempProxyConvos1 = append(tempProxyConvos1, elem)
			}
		}

		err = post.UpdateProxyMsg(client.Eclient, proxyid, "Conversations", tempProxyConvos1)

		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(hit.Id + "failed to be removed from proxy 1")
			log.Println(err)
			return
		}

		tempProxyConvos2 := []types.ConversationState{}
		for _, elem := range proxy2.Conversations {
			if elem.ConvoID != hit.Id {
				tempProxyConvos2 = append(tempProxyConvos2, elem)
			}
		}

		err = post.UpdateProxyMsg(client.Eclient, proxyid2, "Conversations", tempProxyConvos2)

		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(hit.Id + "failed to be removed from proxy 2")
			log.Println(err)
			return
		}

		err = globals.DeleteByID(client.Eclient, hit.Id, "convo")
		if err != nil {
			fmt.Println(hit.Id + "failed to be deleted")
			fmt.Println(err)
		} else {
			fmt.Println("number of chats deleted = " + strconv.Itoa(i))
		}
	}

	// usr, err := get.UserByID(client.Eclient, "7v5wyWgBN3Vvtvdi4pWH")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(usr.FirstName + usr.LastName)

}
