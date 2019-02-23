package main

import (
	"fmt"

	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
	"github.com/sea350/ustart_go/types"
)

func main() {

	//"8v5xyWgBN3VvtvdiWpXP" mins doc id
	//v4e02gBN3VvtvdiDZYs tarek doc id
	//7v5wyWgBN3Vvtvdi4pWH steven doc id

	id, err := get.IDByUsername(client.Eclient, "nevets")
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
	fmt.Println("User msg proxies: " + proxyid)

	proxy, err := getChat.ProxyMsgByID(client.Eclient, proxyid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Printing cached convos: ")

	var tempArr []types.ConversationState
	fmt.Println("len before")
	fmt.Println(len(proxy.Conversations))
	for i := range proxy.Conversations {
		if proxy.Conversations[i].ConvoID == "Ff60-GgBN3VvtvdiHZ2B" || proxy.Conversations[i].ConvoID == "Ev6z-GgBN3Vvtvdi9Z3Q" {
			continue
		}
		tempArr = append(tempArr, proxy.Conversations[i])
	}

	err = postChat.UpdateProxyMsg(client.Eclient, proxyid, "Conversations", tempArr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("len after")
	fmt.Println(len(tempArr))

}
