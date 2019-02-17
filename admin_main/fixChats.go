package main

import (
	"context"
	"fmt"

	getUser "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//9v4r-GgBN3VvtvdieZzG
// g_5h42gBN3VvtvdiWZt3

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {
	ctx := context.Background()
	proxyID := "g_5h42gBN3VvtvdiWZt3"
	usrID, _ := getUser.IDByUsername(eclient, "HeatherMT")
	// convoID := "9v4r-GgBN3VvtvdieZzG"

	// fmt.Println("Deleting convo")
	// _, err := eclient.Delete().
	// 	Index(globals.ConvoIndex).
	// 	Type(globals.ConvoType).
	// 	Id(focusID).
	// 	Do(ctx)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	var theProxy = types.ProxyMessages{
		DocID:         usrID,
		Class:         1,
		NumUnread:     0,
		Conversations: nil,
	}

	_, err := eclient.Index().
		Index(globals.ProxyMsgID).
		Type(globals.ProxyMsgType).
		Id(proxyID).
		BodyJson(theProxy).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}
}
