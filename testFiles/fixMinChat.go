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

	id, err := get.IDByUsername(client.Eclient, "min")
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
	for i := range proxy.Conversations {
		if proxy.Conversations[i].ConvoID == "lv5J02gBN3Vvtvdi4JiP" || "7v4i-GgBN3VvtvdiZZzR" == proxy.Conversations[i].ConvoID || "I_5l02gBN3VvtvditZkF" == proxy.Conversations[i].ConvoID || "kf5J02gBN3VvtvdixZgF" == proxy.Conversations[i].ConvoID {
			continue
		}
		tempArr = append(tempArr, proxy.Conversations[i])
	}

	err = postChat.UpdateProxyMsg(client.Eclient, proxyid, "Conversations", tempArr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(tempArr))

}
