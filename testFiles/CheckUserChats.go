package main

import (
	"fmt"

	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
)

func main() {
	fmt.Println("getting user by email")
	id, err := get.IDByUsername(client.Eclient, "th1750")
	if err != nil {
		fmt.Println(err)
		return
	}
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

}
