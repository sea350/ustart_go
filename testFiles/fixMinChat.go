package main

import (
	"context"
	"fmt"

	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
)

func main() {
	ctx := context.Background()

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
	id = "v4e02gBN3VvtvdiDZYs"

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

}
