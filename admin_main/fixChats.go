package main

import (
	"context"
	"fmt"
	"strings"

	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"
	elastic "gopkg.in/olivere/elastic.v5"
)

//9v4r-GgBN3VvtvdieZzG
// g_5h42gBN3VvtvdiWZt3


func main() {
	ctx := context.Background()
	proxyID := g_5h42gBN3VvtvdiWZt3
	focusID := 9v4r-GgBN3VvtvdieZzG
	
	fmt.Println("Deleting convo")
	err = eclient.Delete().
		Index(globals.ConvoIndex).
		Type(globals.ConvoType).
		Id(focusID).
		Do(ctx)
	
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Deleting proxy")
	err = eclient.Delete().
		Index(globals.ProxyMsgIndex).
		Type(globals.ProxyMsgType).
		Id(proxyID).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}
	


}
